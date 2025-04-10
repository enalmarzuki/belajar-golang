/*
 * TodoList RESTful API
 *
 * OpenAPI for TodoList RESTful API
 *
 * API version: 1
 * Contact: testing@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateOrUpdateTodoList struct {
	Name string `json:"name,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
