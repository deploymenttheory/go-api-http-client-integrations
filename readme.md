# Go HTTP Client API Integration Library

This library provides a collection of API integrations utilized by by the [go-api-http-client](https://github.com/deploymenttheory/go-api-http-client) for various applications within the enterprise. The integrations are designed to be flexible, leveraging private functions to meet interface requirements by the http client. This approach allows for support different api specifics, such as authentication methods, request marshalling, header defintions and so on while maintaining a consistent interface for the http client.

## Overview

An API integration in this library typically includes the following components:

1. **Authentication Methods**:
   - **OAuth2.0**: Implements client credentials flow.
   - **Basic Authentication**: Uses username and password.
   - **Token Management**: Methods to handle token retrieval, expiration checks, and refreshing tokens.

2. **Request Preparation**:
   - **Header Management**: Setting required headers such as `Accept`, `Content-Type`, and `Authorization`.
   - **Header Exceptions**: Handling specific headers for certain requests.
   - **Request Body Marshalling**: Encoding the request body as JSON, XML or multipart form data.

3. **Integration Builders**:
    - **Initialization Functions**: Methods to initialize the integration with specific authentication methods, such as OAuth2.0 and Basic Authentication.
    - **Configuration Parameters**: Contextual parameters necessary for setting up the integration. These parameters vary based on the API and the authentication method being used. Examples include:
    - **Microsoft MS Graph**:
       - `clientId`, `clientSecret`, `tenantID`, `bufferPeriod`
    - **Jamf Pro**:
       - `clientId`, `clientSecret`, `username`, `password`, `jamfBaseDomain`, `bufferPeriod`

4. **Shared Utilities**: Common utility functions that are common practise across multiple API integrations.

dev/bob
