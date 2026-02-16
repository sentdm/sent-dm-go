# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListItem">ContactListItem</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListResponse">ContactListResponse</a>

Methods:

- <code title="get /v2/contacts">client.Contacts.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListParams">ContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListResponse">ContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/contacts/phone">client.Contacts.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactService.GetByPhone">GetByPhone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactGetByPhoneParams">ContactGetByPhoneParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListItem">ContactListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/contacts/id">client.Contacts.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactService.GetID">GetID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactGetIDParams">ContactGetIDParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#ContactListItem">ContactListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageGetResponse">MessageGetResponse</a>

Methods:

- <code title="get /v2/messages/{id}">client.Messages.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageGetResponse">MessageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/messages/quick-message">client.Messages.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageService.SendQuickMessage">SendQuickMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageSendQuickMessageParams">MessageSendQuickMessageParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v2/messages/contact">client.Messages.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageService.SendToContact">SendToContact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageSendToContactParams">MessageSendToContactParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v2/messages/phone">client.Messages.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageService.SendToPhone">SendToPhone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#MessageSendToPhoneParams">MessageSendToPhoneParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateBodyContentParam">TemplateBodyContentParam</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateDefinitionParam">TemplateDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateVariableParam">TemplateVariableParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateBodyContent">TemplateBodyContent</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateDefinition">TemplateDefinition</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateVariable">TemplateVariable</a>
- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateListResponse">TemplateListResponse</a>

Methods:

- <code title="post /v2/templates">client.Templates.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateNewParams">TemplateNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/templates/{id}">client.Templates.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateResponseV2">TemplateResponseV2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/templates">client.Templates.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateListParams">TemplateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateListResponse">TemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/templates/{id}">client.Templates.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#TemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# NumberLookup

Response Types:

- <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#NumberLookupGetResponse">NumberLookupGetResponse</a>

Methods:

- <code title="get /v2/number-lookup">client.NumberLookup.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#NumberLookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#NumberLookupGetParams">NumberLookupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go">sentdm</a>.<a href="https://pkg.go.dev/github.com/sentdm/sent-dm-go#NumberLookupGetResponse">NumberLookupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
