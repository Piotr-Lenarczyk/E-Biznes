package controllers

import play.api.mvc._
import play.api.libs.json._

import javax.inject._
import scala.collection.mutable

@Singleton
class BasketController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  // Sample in-memory data store
  private val baskets = mutable.ListBuffer(
    Basket(Some(1), List(1, 2)),  // Basket with products 1 and 2
    Basket(Some(2), List(3))      // Basket with product 3
  )

  implicit val basketFormat: OFormat[Basket] = Json.format[Basket]

  // List all baskets (GET /baskets)
  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(baskets))
  }

  // Get basket by ID (GET /baskets/:id)
  def getById(id: Int): Action[AnyContent] = Action {
    baskets.find(_.id.contains(id)) match {
      case Some(basket) => Ok(Json.toJson(basket))
      case None         => NotFound(Json.obj("error" -> "Basket not found"))
    }
  }

  // Create a new basket (POST /baskets)
  def create: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Basket] match {
      case JsSuccess(basket, _) =>
        val newId = baskets.flatMap(_.id).maxOption.getOrElse(0) + 1
        val newBasket = basket.copy(id = Some(newId))
        baskets += newBasket
        Created(Json.toJson(newBasket))

      case JsError(errors) =>
        BadRequest(Json.obj(
          "error" -> "Invalid JSON",
          "details" -> JsError.toJson(errors)
        ))
    }
  }

  // Update an existing basket (PUT /baskets/:id)
  def update(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Basket] match {
      case JsSuccess(updatedBasket, _) =>
        baskets.indexWhere(_.id.contains(id)) match {
          case -1 => NotFound(Json.obj("error" -> "Basket not found"))
          case index =>
            baskets.update(index, updatedBasket.copy(id = Some(id)))
            Ok(Json.toJson(updatedBasket.copy(id = Some(id))))
        }

      case JsError(errors) =>
        BadRequest(Json.obj(
          "error" -> "Invalid JSON",
          "details" -> JsError.toJson(errors)
        ))
    }
  }

  // Delete a basket (DELETE /baskets/:id)
  def delete(id: Int): Action[AnyContent] = Action {
    baskets.indexWhere(_.id.contains(id)) match {
      case -1 => NotFound(Json.obj("error" -> "Basket not found"))
      case index =>
        baskets.remove(index)
        NoContent
    }
  }
}

// Basket case class
case class Basket(id: Option[Int], productIds: List[Int])

object Basket {
  implicit val basketFormat: Format[Basket] = Json.format[Basket]
}
