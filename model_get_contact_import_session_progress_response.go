/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetContactImportSessionProgressResponse struct {
	// Session status: * **1** - if session has been initialized but not yet started; * **3** - if session is being processed; * **4** - if session has errors; * **5** - if session completed successfully. 
	Status int32 `json:"status"`
	// How many contacts have been imported?
	Processed int32 `json:"processed"`
}
