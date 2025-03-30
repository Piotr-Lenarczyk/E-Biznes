package controllers

import play.api.mvc._
import play.api.libs.json._

import javax.inject._
import scala.collection.mutable

@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  // Sample in-memory data store
  private val categories = mutable.ListBuffer(
    Category(Some(1), "Electronics"),
    Category(Some(2), "Clothing"),
    Category(Some(3), "Books")
  )

  implicit val categoryFormat: OFormat[Category] = Json.format[Category]

  // List all categories (GET /categories)
  def getAll: Action[AnyContent] = Action {
    Ok(Json.toJson(categories))
  }

  // Get category by ID (GET /categories/:id)
  def getById(id: Int): Action[AnyContent] = Action {
    categories.find(_.id == Some(id)) match {
      case Some(category) => Ok(Json.toJson(category))
      case None           => NotFound(Json.obj("error" -> "Category not found"))
    }
  }

  // Create a new category (POST /categories)
  def create: Action[JsValue] = Action(parse.json) { request =>
  request.body.validate[Category] match {
    case JsSuccess(category, _) =>
      val newId = categories.map(_.id.getOrElse(0)).maxOption.getOrElse(0) + 1
      val newCategory = category.copy(id = Some(newId))
      categories += newCategory
      Created(Json.toJson(newCategory))

    case JsError(errors) =>
      BadRequest(Json.obj(
        "error" -> "Invalid JSON",
        "details" -> JsError.toJson(errors)
      ))
  }
}

  // Update category (PUT /categories/:id)
  def update(id: Int): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category] match {
      case JsSuccess(updatedCategory, _) =>
        categories.indexWhere(_.id == Some(id)) match {
          case -1 => NotFound(Json.obj("error" -> "Category not found"))
          case index =>
            categories.update(index, updatedCategory.copy(id = Some(id)))
            Ok(Json.toJson(updatedCategory.copy(id = Some(id))))
        }

      case JsError(errors) =>
        BadRequest(Json.obj(
          "error" -> "Invalid JSON",
          "details" -> JsError.toJson(errors)
        ))
    }
  }

  // Delete category (DELETE /categories/:id)
  def delete(id: Int): Action[AnyContent] = Action {
    categories.indexWhere(_.id == Some(id)) match {
      case -1 => NotFound(Json.obj("error" -> "Category not found"))
      case index =>
        categories.remove(index)
        NoContent
    }
  }
}

// Category case class
case class Category(id: Option[Int], name: String)

object Category {
  implicit val categoryFormat: OFormat[Category] = Json.format[Category]
}
