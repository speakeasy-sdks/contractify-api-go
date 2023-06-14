# ContractifyProduction SDK

## Overview

Contractify Public API: This is the public API for integrating with Contractify

# Introduction

The Contractify API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.

# Authentication

The Contractify API uses the OAuth2 protocol for authentication. To be able to authenticate against the API, you need to register an application. A registered app is assigned a unique client ID and client secret which will be used in the OAuth flow. The client secret should not be shared. Managing your OAuth applications can be done by navigating to [your developer settings](/client/profile#developer). You can mail us at [api@contractify.be](mailto:api@contractify.be) to receive fresh API credentials.

The [Authorization Code Flow](https://developer.okta.com/blog/2018/04/10/oauth-authorization-code-grant-type) can be used to receive API access with a manual confirmation step. The authorization URL and Token URL can be found in the [Authentication explorer](#auth).

Access tokens expire after 1 month, but can be refreshed using the refresh token that is issued together with your access token. This can be done by calling our Token URL with a refresh_token grant_type. You can find more information about the refresh token flow [here](https://www.oauth.com/oauth2-servers/access-tokens/refreshing-access-tokens/).

# Rate limiting

The Contractify API employs several safeguards against bursts of incoming traffic to help maximize its stability. Clients that send many requests in quick succession may see error responses that show up as status code 429. The API allows up to 1000 requests per minute per client. Every request includes the amount of available and remaining requests.

| header                | explanation                                               |
|-----------------------|-----------------------------------------------------------|
| X-RateLimit-Limit     | The number of requests that are allowed                   |
| X-RateLimit-Remaining | The number of requests that are still available           |

When hitting the rate limit, we provide a couple of additional headers that will help you implement a retry mechanism.

| header                | explanation                                               |
|-----------------------|-----------------------------------------------------------|
| Retry-After           | The number of seconds before the rate limit will be reset |
| X-RateLimit-Reset     | The timestamp when the rate limit will be reset           |

# Pagination

Requests that return multiple items will be paginated to 100 items by default. You can specify further pages with the page parameter.

Note that page numbering is 1-based and that omitting the page parameter will return the first page.

# Versioning

When backwards-incompatible changes will be made to the API, a new, dated version will be released. The current latest version is *2022-01-18*. All requests use your applications configured version unless you override it by specifying another date in the X-Api-Version. This can help you upgrade your integration to a newer API version gradually.

# Partial updates

To avoid potential confusion between omitted fields and null values, we do not allow for partial updates. All fields should be provided when doing create and update endpoints, but non-required fields can be passed with a null value.

# Durations

There are multiple places in the API where we allow for sending and receiving durations (sometimes also called periods or intervals). For these durations, we use the ISO 8601 duration format, with one time identifier, and only one of the following time elements: Y, M, W, D. An example would be `P3M` for a period of 3 months.

### Available Operations

