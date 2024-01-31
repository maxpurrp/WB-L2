package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Паттерн цепочка вызовов - поведенческий шаблон с помощью которого можно передавать запросы
			последовательно по цепочке обработчиков. Каждый обработчик решает, может ли он
			обработать запрос, и либо обрабатывает его, либо передает следующему обработчику в цепочке.
Плюсы - 1. Уменьшение связанности кода: Каждый обработчик концентрируется на конкретной обязанности, и
		   изменения в одном обработчике не влияют на другие. Это способствует легкости
		   сопровождения и расширения кода.
		2. Гибкость и расширяемость: Цепочка обработчиков позволяет гибко настраивать порядок и условия
		   обработки запросов. Добавление новых обработчиков или изменение порядка существующих не требует
		   изменений в клиентском коде.

Минусы - 1. Гарантия обработки: Не всегда гарантируется, что запрос будет обработан каким-либо обработчиком в цепочке.
			В случае отсутствия подходящего обработчика запрос может быть проигнорирован или зациклен.
		 2. Влияние на производительность: Передача запроса через цепочку обработчиков может повлиять
		    на производительность из-за дополнительных проверок и вызовов методов.
*/

// handler interface defines the common methods for all handlers
type Handler interface {
	HandleRequest(request *HttpRequest) string
	SetNextHandler(handler Handler)
}

// concreteHandler for authentication
type AuthenticationHandler struct {
	nextHandler Handler
}

// handleRequest checks if the request is authenticated
func (h *AuthenticationHandler) HandleRequest(request *HttpRequest) string {
	if request.Authenticated {
		return "Authentication passed. "
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Authentication failed. "
}

// wetNextHandler sets the next handler in the chain
func (h *AuthenticationHandler) SetNextHandler(handler Handler) {
	h.nextHandler = handler
}

// concreteHandler for authorization
type AuthorizationHandler struct {
	nextHandler Handler
}

// handleRequest checks if the request has the required permissions
func (h *AuthorizationHandler) HandleRequest(request *HttpRequest) string {
	if request.HasPermission {
		return "Authorization passed. "
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Authorization failed. "
}

// setNextHandler sets the next handler in the chain
func (h *AuthorizationHandler) SetNextHandler(handler Handler) {
	h.nextHandler = handler
}

// httpRequest represents a sample HTTP request
type HttpRequest struct {
	Authenticated bool
	HasPermission bool
}

// func main() {
// 	// create a sample HTTP request
// 	request := &HttpRequest{Authenticated: true, HasPermission: false}

// 	// create handlers.
// 	authenticationHandler := &AuthenticationHandler{}
// 	authorizationHandler := &AuthorizationHandler{}

// 	// set up the chain of responsibility
// 	authenticationHandler.SetNextHandler(authorizationHandler)

// 	// client initiates a request.
// 	result := authenticationHandler.HandleRequest(request)
// 	fmt.Println("Result:", result)
// }
