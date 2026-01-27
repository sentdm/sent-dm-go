# Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileListTemplatesResponse">ProfileListTemplatesResponse</a>

Methods:

- <code title="get /v3/profiles/{profileId}/templates">client.Profiles.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileService.ListTemplates">ListTemplates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileListTemplatesParams">ProfileListTemplatesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileListTemplatesResponse">ProfileListTemplatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v3/profiles/{profileId}/messages">client.Profiles.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileSendMessageParams">ProfileSendMessageParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Users

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InviteUserRequestParam">InviteUserRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#UpdateUserRoleRequestParam">UpdateUserRoleRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#BaseDto">BaseDto</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InviteUserResponse">InviteUserResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#UserListResponse">UserListResponse</a>

Methods:

- <code title="get /v3/profiles/{profileId}/users/{userId}">client.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserGetParams">ProfileUserGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/profiles/{profileId}/users">client.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserListParams">ProfileUserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/profiles/{profileId}/users/{userId}">client.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserDeleteParams">ProfileUserDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/profiles/{profileId}/users/invite">client.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserInviteParams">ProfileUserInviteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InviteUserResponse">InviteUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/profiles/{profileId}/users/{userId}/role">client.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileUserUpdateRoleParams">ProfileUserUpdateRoleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListItemProfile">ContactListItemProfile</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactListResponse">ProfileContactListResponse</a>

Methods:

- <code title="get /v3/profiles/{profileId}/contacts/{contactId}">client.Profiles.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, contactID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactGetParams">ProfileContactGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ContactListItemProfile">ContactListItemProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/profiles/{profileId}/contacts">client.Profiles.Contacts.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, profileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactListParams">ProfileContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#ProfileContactListResponse">ProfileContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InvitationDetails">InvitationDetails</a>

Methods:

- <code title="get /v3/organizations/{customerId}/profiles/{profileId}/users/invitations/{token}">client.Organizations.Profiles.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationProfileUserService.GetInvitationDetails">GetInvitationDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationProfileUserGetInvitationDetailsParams">OrganizationProfileUserGetInvitationDetailsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InvitationDetails">InvitationDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerResponse">OrganizationUserListByCustomerResponse</a>

Methods:

- <code title="get /v3/organizations/{orgId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserGetParams">OrganizationUserGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/organizations/{orgId}/users">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListParams">OrganizationUserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/organizations/{orgId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserDeleteParams">OrganizationUserDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v2/organizations/{customerId}/users">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.NewOrInvite">NewOrInvite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserNewOrInviteParams">OrganizationUserNewOrInviteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.DeleteByCustomer">DeleteByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserDeleteByCustomerParams">OrganizationUserDeleteByCustomerParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v3/organizations/{orgId}/users/invite">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserInviteParams">OrganizationUserInviteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InviteUserResponse">InviteUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/organizations/{customerId}/users">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.ListByCustomer">ListByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerParams">OrganizationUserListByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserListByCustomerResponse">OrganizationUserListByCustomerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.GetByCustomer">GetByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserGetByCustomerParams">OrganizationUserGetByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/organizations/{customerId}/users/invitations/{token}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.GetInvitationDetails">GetInvitationDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, token <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserGetInvitationDetailsParams">OrganizationUserGetInvitationDetailsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#InvitationDetails">InvitationDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v3/organizations/{orgId}/users/{userId}/role">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserUpdateRoleParams">OrganizationUserUpdateRoleParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /v2/organizations/{customerId}/users/{userId}">client.Organizations.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserService.UpdateRoleByCustomer">UpdateRoleByCustomer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#OrganizationUserUpdateRoleByCustomerParams">OrganizationUserUpdateRoleByCustomerParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#CustomerUserDto">CustomerUserDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Healthcheck

Methods:

- <code title="get /healthcheck">client.Healthcheck.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#HealthcheckService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Health

Methods:

- <code title="get /health/live">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#HealthService.CheckLive">CheckLive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /health/ready">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/sent-dm-go#HealthService.CheckReady">CheckReady</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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
