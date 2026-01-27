# Profiles

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#BaseDto">BaseDto</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>

## Contacts

# Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationListAuthenticatedUserOrganizationsResponse">OrganizationListAuthenticatedUserOrganizationsResponse</a>

Methods:

- <code title="get /v2/organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationService.ListAuthenticatedUserOrganizations">ListAuthenticatedUserOrganizations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationListAuthenticatedUserOrganizationsResponse">OrganizationListAuthenticatedUserOrganizationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileSummary">ProfileSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationProfileListResponse">OrganizationProfileListResponse</a>

Methods:

- <code title="get /v2/organizations/{orgId}/profiles">client.Organizations.Profiles.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationProfileListResponse">OrganizationProfileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerResponse">OrganizationUserListByCustomerResponse</a>

Methods:

- <code title="post /v2/organizations/{customerId}/users">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.NewOrInvite">NewOrInvite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserNewOrInviteParams">OrganizationUserNewOrInviteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.DeleteByCustomer">DeleteByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserDeleteByCustomerParams">OrganizationUserDeleteByCustomerParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v2/organizations/{customerId}/users">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.ListByCustomer">ListByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerParams">OrganizationUserListByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerResponse">OrganizationUserListByCustomerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.GetByCustomer">GetByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserGetByCustomerParams">OrganizationUserGetByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.UpdateRoleByCustomer">UpdateRoleByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserUpdateRoleByCustomerParams">OrganizationUserUpdateRoleByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Healthcheck

# Health

# Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateBodyContentParam">TemplateBodyContentParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateDefinitionParam">TemplateDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateVariableParam">TemplateVariableParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateBodyContent">TemplateBodyContent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateDefinition">TemplateDefinition</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateVariable">TemplateVariable</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateListResponse">TemplateListResponse</a>

Methods:

- <code title="post /v2/templates">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateNewParams">TemplateNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/templates/{id}">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/templates">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateListParams">TemplateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateListResponse">TemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/templates/{id}">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#TemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListItemV2">ContactListItemV2</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListResponse">ContactListResponse</a>

Methods:

- <code title="get /v2/contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListParams">ContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListResponse">ContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/contacts/phone">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactService.GetByPhone">GetByPhone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactGetByPhoneParams">ContactGetByPhoneParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListItemV2">ContactListItemV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/contacts/id">client.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactService.GetID">GetID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactGetIDParams">ContactGetIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListItemV2">ContactListItemV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageGetResponse">MessageGetResponse</a>

Methods:

- <code title="get /v2/messages/{id}">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageGetResponse">MessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/messages/quick-message">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageService.SendQuickMessage">SendQuickMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageSendQuickMessageParams">MessageSendQuickMessageParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v2/messages/contact">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageService.SendToContact">SendToContact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageSendToContactParams">MessageSendToContactParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v2/messages/phone">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageService.SendToPhone">SendToPhone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#MessageSendToPhoneParams">MessageSendToPhoneParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# NumberLookup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#NumberLookupGetResponse">NumberLookupGetResponse</a>

Methods:

- <code title="get /v2/number-lookup">client.NumberLookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#NumberLookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#NumberLookupGetParams">NumberLookupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#NumberLookupGetResponse">NumberLookupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
