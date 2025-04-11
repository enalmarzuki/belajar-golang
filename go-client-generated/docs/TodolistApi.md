# {{classname}}

All URIs are relative to *https://{environment}.enalmarzuki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TodolistGet**](TodolistApi.md#TodolistGet) | **Get** /todolist | Get all todolist
[**TodolistPost**](TodolistApi.md#TodolistPost) | **Post** /todolist | Create new todolist
[**TodolistTodoListIdDelete**](TodolistApi.md#TodolistTodoListIdDelete) | **Delete** /todolist/{todoListId} | Delete existing todolist
[**TodolistTodoListIdPut**](TodolistApi.md#TodolistTodoListIdPut) | **Put** /todolist/{todoListId} | Update existing todolist

# **TodolistGet**
> []TodoList TodolistGet(ctx, optional)
Get all todolist

Get all active todolist by default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TodolistApiTodolistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TodolistApiTodolistGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDone** | **optional.Bool**| Is include done todolist | [default to false]
 **name** | **optional.String**| Filter todolist by name | 

### Return type

[**[]TodoList**](array.md)

### Authorization

[TodoListAuth](../README.md#TodoListAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistPost**
> TodoList TodolistPost(ctx, body)
Create new todolist

Create new todolist to database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateTodoList**](CreateOrUpdateTodoList.md)|  | 

### Return type

[**TodoList**](TodoList.md)

### Authorization

[TodoListAuth](../README.md#TodoListAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistTodoListIdDelete**
> InlineResponse200 TodolistTodoListIdDelete(ctx, todoListId)
Delete existing todolist

Delete existing todolist in database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todoListId** | **string**| Todolist id for update | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[TodoListAuth](../README.md#TodoListAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistTodoListIdPut**
> TodoList TodolistTodoListIdPut(ctx, body, todoListId)
Update existing todolist

Update existing todolist in database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TodolistTodoListIdBody**](TodolistTodoListIdBody.md)|  | 
  **todoListId** | **string**| Todolist id for update | 

### Return type

[**TodoList**](TodoList.md)

### Authorization

[TodoListAuth](../README.md#TodoListAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

