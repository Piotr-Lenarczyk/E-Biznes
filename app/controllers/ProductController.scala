package controllers

import play.api.mvc._
import javax.inject._

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  def placeholder: Action[AnyContent] = Action {
    Ok("ProductController is temporarily unavailable.")
  }
}
