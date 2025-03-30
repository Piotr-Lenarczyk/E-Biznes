package controllers

import play.api.mvc._
import play.api.libs.json._

import javax.inject._
import scala.collection.mutable

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  // Sample in-memory data store
  private val products = mutable.ListBuffer(
    Product(Some(1), "Laptop", 1200.50),
    Product(Some(2), "Phone", 800.00),
    Product(Some(3), "Tablet", 500.99)
  )

  implicit val productFormat: OFormat[Product] = Json.format[Product]

  // List all products (GET /products)
  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(products))
  }

  // Get product by ID (GET /products/:id)
  def getById(id: Int): Action[AnyContent] = Action {
    products.find(_.id.contains(id)) match {
      case Some(product) => Ok(Json.toJson(product))
      case None          => NotFound(Json.obj("error" -> "Product not found"))
    }
  }

  // Create a new product (POST /products)
  def create: Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product] match {
      case JsSuccess(product, _) =>
        val newId = products.flatMap(_.id).maxOption.getOrElse(0) + 1
        val newProduct = product.copy(id = Some(newId))
        products += newProduct
        Created(Json.toJson(newProduct))

      case JsError(errors) =>
        BadRequest(Json.obj(
          "error" -> "Invalid JSON",
          "details" -> JsError.toJson(errors)
        ))
    }
  }

  // Update an existing product (PUT /products/:id)
  def update(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product] match {
      case JsSuccess(updatedProduct, _) =>
        products.indexWhere(_.id.contains(id)) match {
          case -1 => NotFound(Json.obj("error" -> "Product not found"))
          case index =>
            products.update(index, updatedProduct.copy(id = Some(id)))
            Ok(Json.toJson(updatedProduct.copy(id = Some(id))))
        }

      case JsError(errors) =>
        BadRequest(Json.obj(
          "error" -> "Invalid JSON",
          "details" -> JsError.toJson(errors)
        ))
    }
  }

  // Delete a product (DELETE /products/:id)
  def delete(id: Int): Action[AnyContent] = Action {
    products.indexWhere(_.id.contains(id)) match {
      case -1 => NotFound(Json.obj("error" -> "Product not found"))
      case index =>
        products.remove(index)
        NoContent
    }
  }
}

// Product case class
case class Product(id: Option[Int], name: String, price: Double)

object Product {
  implicit val productFormat: Format[Product] = Json.format[Product]
}
