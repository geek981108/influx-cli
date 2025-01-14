/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	_context "context"
	_fmt "fmt"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type SetupApi interface {

	/*
	 * GetSetup Check if database has default user, org, bucket
	 * Returns `true` if no default user, organization, or bucket has been created.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetSetupRequest
	 */
	GetSetup(ctx _context.Context) ApiGetSetupRequest

	/*
	 * GetSetupExecute executes the request
	 * @return InlineResponse200
	 */
	GetSetupExecute(r ApiGetSetupRequest) (InlineResponse200, error)

	/*
	 * PostSetup Set up initial user, org and bucket
	 * Post an onboarding request to set up initial user, org and bucket.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostSetupRequest
	 */
	PostSetup(ctx _context.Context) ApiPostSetupRequest

	/*
	 * PostSetupExecute executes the request
	 * @return OnboardingResponse
	 */
	PostSetupExecute(r ApiPostSetupRequest) (OnboardingResponse, error)

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on OSS
	// servers.
	OnlyOSS() SetupApi

	// Sets additional descriptive text in the error message if any request in
	// this API fails, indicating that it is intended to be used only on cloud
	// servers.
	OnlyCloud() SetupApi
}

// SetupApiService SetupApi service
type SetupApiService service

func (a *SetupApiService) OnlyOSS() SetupApi {
	a.isOnlyOSS = true
	return a
}

func (a *SetupApiService) OnlyCloud() SetupApi {
	a.isOnlyCloud = true
	return a
}

type ApiGetSetupRequest struct {
	ctx          _context.Context
	ApiService   SetupApi
	zapTraceSpan *string
}

func (r ApiGetSetupRequest) ZapTraceSpan(zapTraceSpan string) ApiGetSetupRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiGetSetupRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiGetSetupRequest) Execute() (InlineResponse200, error) {
	return r.ApiService.GetSetupExecute(r)
}

/*
 * GetSetup Check if database has default user, org, bucket
 * Returns `true` if no default user, organization, or bucket has been created.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetSetupRequest
 */
func (a *SetupApiService) GetSetup(ctx _context.Context) ApiGetSetupRequest {
	return ApiGetSetupRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *SetupApiService) GetSetupExecute(r ApiGetSetupRequest) (InlineResponse200, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetupApiService.GetSetup")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/setup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type ApiPostSetupRequest struct {
	ctx               _context.Context
	ApiService        SetupApi
	onboardingRequest *OnboardingRequest
	zapTraceSpan      *string
}

func (r ApiPostSetupRequest) OnboardingRequest(onboardingRequest OnboardingRequest) ApiPostSetupRequest {
	r.onboardingRequest = &onboardingRequest
	return r
}
func (r ApiPostSetupRequest) GetOnboardingRequest() *OnboardingRequest {
	return r.onboardingRequest
}

func (r ApiPostSetupRequest) ZapTraceSpan(zapTraceSpan string) ApiPostSetupRequest {
	r.zapTraceSpan = &zapTraceSpan
	return r
}
func (r ApiPostSetupRequest) GetZapTraceSpan() *string {
	return r.zapTraceSpan
}

func (r ApiPostSetupRequest) Execute() (OnboardingResponse, error) {
	return r.ApiService.PostSetupExecute(r)
}

/*
 * PostSetup Set up initial user, org and bucket
 * Post an onboarding request to set up initial user, org and bucket.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostSetupRequest
 */
func (a *SetupApiService) PostSetup(ctx _context.Context) ApiPostSetupRequest {
	return ApiPostSetupRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return OnboardingResponse
 */
func (a *SetupApiService) PostSetupExecute(r ApiPostSetupRequest) (OnboardingResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OnboardingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SetupApiService.PostSetup")
	if err != nil {
		return localVarReturnValue, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/setup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.onboardingRequest == nil {
		return localVarReturnValue, reportError("onboardingRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.zapTraceSpan != nil {
		localVarHeaderParams["Zap-Trace-Span"] = parameterToString(*r.zapTraceSpan, "")
	}
	// body params
	localVarPostBody = r.onboardingRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	var errorPrefix string
	if a.isOnlyOSS {
		errorPrefix = "InfluxDB OSS-only command failed: "
	} else if a.isOnlyCloud {
		errorPrefix = "InfluxDB Cloud-only command failed: "
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		body, err := GunzipIfNeeded(localVarHTTPResponse)
		if err != nil {
			body.Close()
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		localVarBody, err := _io.ReadAll(body)
		body.Close()
		if err != nil {
			return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
		}
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, localVarHTTPResponse.Status),
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = _fmt.Sprintf("%s: %s", newErr.Error(), err.Error())
			return localVarReturnValue, newErr
		}
		v.SetMessage(_fmt.Sprintf("%s: %s", newErr.Error(), v.GetMessage()))
		newErr.model = &v
		return localVarReturnValue, newErr
	}

	body, err := GunzipIfNeeded(localVarHTTPResponse)
	if err != nil {
		body.Close()
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	localVarBody, err := _io.ReadAll(body)
	body.Close()
	if err != nil {
		return localVarReturnValue, _fmt.Errorf("%s%w", errorPrefix, err)
	}
	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: _fmt.Sprintf("%s%s", errorPrefix, err.Error()),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}
