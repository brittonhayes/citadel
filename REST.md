---
title: Citadel service v1.0
language_tabs:
  - golang: Golang
language_clients:
  - golang: ""
toc_footers: []
includes: []
search: false
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="citadel-service">Citadel service v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

HTTP service for working with Security Operations Center resources

Base URLs:

* <a href="http://localhost:8000/citadel">http://localhost:8000/citadel</a>

<h1 id="citadel-service-incidents">incidents</h1>

## list all incidents

<a id="opIdincidents#list all"></a>

> Code samples

`GET /incidents`

<h3 id="list-all-incidents-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|limit|query|integer|false|Limit the number of results|

> Example responses

> 200 Response

```json
[
  {
    "Permissions": "analyst",
    "affected_customers": [
      "Quis ipsam.",
      "Voluptatem eum.",
      "Vero aliquam quisquam veritatis quia quis delectus."
    ],
    "created_at": "1994-12-30T14:30:47Z",
    "date": "1996-10-26",
    "date_closed": "1993-08-14",
    "id": 12043849106673500000,
    "responsible_party": "Recusandae quae nobis.",
    "root_cause": "Aut voluptatem deserunt non dolor omnis.",
    "scope": "Qui assumenda vitae reprehenderit dolores.",
    "severity": 0,
    "slack_channel": "Aliquid dolore harum.",
    "summary": "Ipsam id praesentium blanditiis adipisci.",
    "title": "Tempora velit sed minima distinctio velit.",
    "updated_at": "2004-11-14T00:10:42Z"
  },
  {
    "Permissions": "analyst",
    "affected_customers": [
      "Quis ipsam.",
      "Voluptatem eum.",
      "Vero aliquam quisquam veritatis quia quis delectus."
    ],
    "created_at": "1994-12-30T14:30:47Z",
    "date": "1996-10-26",
    "date_closed": "1993-08-14",
    "id": 12043849106673500000,
    "responsible_party": "Recusandae quae nobis.",
    "root_cause": "Aut voluptatem deserunt non dolor omnis.",
    "scope": "Qui assumenda vitae reprehenderit dolores.",
    "severity": 0,
    "slack_channel": "Aliquid dolore harum.",
    "summary": "Ipsam id praesentium blanditiis adipisci.",
    "title": "Tempora velit sed minima distinctio velit.",
    "updated_at": "2004-11-14T00:10:42Z"
  },
  {
    "Permissions": "analyst",
    "affected_customers": [
      "Quis ipsam.",
      "Voluptatem eum.",
      "Vero aliquam quisquam veritatis quia quis delectus."
    ],
    "created_at": "1994-12-30T14:30:47Z",
    "date": "1996-10-26",
    "date_closed": "1993-08-14",
    "id": 12043849106673500000,
    "responsible_party": "Recusandae quae nobis.",
    "root_cause": "Aut voluptatem deserunt non dolor omnis.",
    "scope": "Qui assumenda vitae reprehenderit dolores.",
    "severity": 0,
    "slack_channel": "Aliquid dolore harum.",
    "summary": "Ipsam id praesentium blanditiis adipisci.",
    "title": "Tempora velit sed minima distinctio velit.",
    "updated_at": "2004-11-14T00:10:42Z"
  }
]
```

<h3 id="list-all-incidents-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK response.|Inline|

<h3 id="list-all-incidents-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Incident](#schemaincident)]|false|none|none|
|» Permissions|string|false|none|Permissions associated with incident|
|» affected_customers|[string]|false|none|A list of the affected customers|
|» created_at|string(date-time)|false|none|When the incident was submitted|
|» date|string(date)|false|none|Date the incident occurred|
|» date_closed|string(date)|false|none|Date the incident occurred|
|» id|integer|false|none|Unique ID of the incident|
|» responsible_party|string|false|none|What group or individual caused the initial incident|
|» root_cause|string|false|none|The original cause of the incident|
|» scope|string|false|none|The scope of impact of this incident|
|» severity|integer|false|none|The severity of the incident|
|» slack_channel|string|false|none|The slack channel for incident discussions|
|» summary|string|false|none|The detailed description of the incident|
|» title|string|false|none|The short title of the incident|
|» updated_at|string(date-time)|false|none|When the incident was last updated|

#### Enumerated Values

|Property|Value|
|---|---|
|Permissions|admin|
|Permissions|analyst|
|Permissions|guest|
|severity|0|
|severity|1|
|severity|2|
|severity|3|
|severity|4|
|severity|5|

<aside class="success">
This operation does not require authentication
</aside>

## find incidents

<a id="opIdincidents#find"></a>

> Code samples

`GET /incidents/{id}`

<h3 id="find-incidents-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|integer|true|Unique ID of the incident|

> Example responses

> 200 Response

```json
{
  "Permissions": "admin",
  "affected_customers": [
    "Aut facere.",
    "Perspiciatis enim.",
    "Dignissimos magnam asperiores ullam dolorem sit ex."
  ],
  "created_at": "1995-05-02T06:56:39Z",
  "date": "1977-02-11",
  "date_closed": "1982-10-28",
  "id": 11941844764433738000,
  "responsible_party": "Et illo ex aliquid veniam nam facere.",
  "root_cause": "Blanditiis autem.",
  "scope": "Libero voluptatum provident esse odio.",
  "severity": 0,
  "slack_channel": "Et ipsum et eum.",
  "summary": "Ratione porro esse suscipit et.",
  "title": "Totam quia.",
  "updated_at": "1972-11-06T07:23:40Z"
}
```

<h3 id="find-incidents-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK response.|[Incident](#schemaincident)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="citadel-service-vulnerabilities">vulnerabilities</h1>

## list vulnerabilities

<a id="opIdvulnerabilities#list"></a>

> Code samples

`GET /vulnerabilities`

List all of the vulnerabilities

<h3 id="list-vulnerabilities-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|limit|query|integer|false|Limit the number of results|

> Example responses

> 200 Response

```json
[
  {
    "cvss_score": 4.7,
    "description": "There is a possible out of bounds write due to a use after free. This could lead to remote code execution with no additional execution privileges needed.",
    "exploitable": false,
    "id": 3819473213214139400,
    "is_patchable": false,
    "is_upgradeable": false,
    "title": "CVE-2020-10"
  },
  {
    "cvss_score": 4.7,
    "description": "There is a possible out of bounds write due to a use after free. This could lead to remote code execution with no additional execution privileges needed.",
    "exploitable": false,
    "id": 3819473213214139400,
    "is_patchable": false,
    "is_upgradeable": false,
    "title": "CVE-2020-10"
  },
  {
    "cvss_score": 4.7,
    "description": "There is a possible out of bounds write due to a use after free. This could lead to remote code execution with no additional execution privileges needed.",
    "exploitable": false,
    "id": 3819473213214139400,
    "is_patchable": false,
    "is_upgradeable": false,
    "title": "CVE-2020-10"
  }
]
```

<h3 id="list-vulnerabilities-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK response.|Inline|

<h3 id="list-vulnerabilities-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Vulnerability](#schemavulnerability)]|false|none|none|
|» cvss_score|number(float)|false|none|Severity score of the vulnerability|
|» description|string|false|none|Description of the vulnerability|
|» exploitable|boolean|false|none|If the vulnerability is exploitable|
|» id|integer|true|none|Unique ID of the vulnerability|
|» is_patchable|boolean|false|none|If the vulnerability is patchable|
|» is_upgradeable|boolean|false|none|If the vulnerability is upgradeable|
|» title|string|false|none|Title of the vulnerability|

<aside class="success">
This operation does not require authentication
</aside>

## submit vulnerabilities

<a id="opIdvulnerabilities#submit"></a>

> Code samples

`POST /vulnerabilities`

> Body parameter

```json
{
  "cvss_score": 0.8337925,
  "description": "Similique est dolor sit corporis excepturi quia.",
  "exploitable": false,
  "is_patchable": true,
  "is_upgradeable": false,
  "title": "Placeat modi debitis."
}
```

<h3 id="submit-vulnerabilities-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[SubmitRequestBody](#schemasubmitrequestbody)|true|none|

> Example responses

> 404 Response

```json
"Vel rerum."
```

<h3 id="submit-vulnerabilities-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created response.|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found response.|[NoMatch](#schemanomatch)|

<aside class="success">
This operation does not require authentication
</aside>

## find vulnerabilities

<a id="opIdvulnerabilities#find"></a>

> Code samples

`GET /vulnerabilities/{id}`

<h3 id="find-vulnerabilities-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|integer|true|Unique ID of the vulnerability|

> Example responses

> 200 Response

```json
{
  "cvss_score": 4.7,
  "description": "There is a possible out of bounds write due to a use after free. This could lead to remote code execution with no additional execution privileges needed.",
  "exploitable": true,
  "id": 11159415604719970000,
  "is_patchable": true,
  "is_upgradeable": false,
  "title": "CVE-2020-10"
}
```

> 404 Response

```json
"Omnis molestiae."
```

<h3 id="find-vulnerabilities-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK response.|[Vulnerability](#schemavulnerability)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found response.|[NoMatch](#schemanomatch)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Incident">Incident</h2>
<!-- backwards compatibility -->
<a id="schemaincident"></a>
<a id="schema_Incident"></a>
<a id="tocSincident"></a>
<a id="tocsincident"></a>

```json
{
  "Permissions": "analyst",
  "affected_customers": [
    "Quo officia ipsum.",
    "Ut ut aut qui voluptatem.",
    "Et at eaque tempore pariatur."
  ],
  "created_at": "1970-11-06T00:06:42Z",
  "date": "1982-06-07",
  "date_closed": "1987-08-18",
  "id": 13669655108109457000,
  "responsible_party": "Fugiat repellendus dolor aperiam ad hic qui.",
  "root_cause": "Deleniti odio est corrupti exercitationem occaecati dicta.",
  "scope": "Numquam rerum et molestiae et culpa.",
  "severity": 2,
  "slack_channel": "Nisi veniam.",
  "summary": "Aliquam quo esse ducimus.",
  "title": "Omnis doloribus.",
  "updated_at": "2012-07-01T19:10:09Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Permissions|string|false|none|Permissions associated with incident|
|affected_customers|[string]|false|none|A list of the affected customers|
|created_at|string(date-time)|false|none|When the incident was submitted|
|date|string(date)|false|none|Date the incident occurred|
|date_closed|string(date)|false|none|Date the incident occurred|
|id|integer|false|none|Unique ID of the incident|
|responsible_party|string|false|none|What group or individual caused the initial incident|
|root_cause|string|false|none|The original cause of the incident|
|scope|string|false|none|The scope of impact of this incident|
|severity|integer|false|none|The severity of the incident|
|slack_channel|string|false|none|The slack channel for incident discussions|
|summary|string|false|none|The detailed description of the incident|
|title|string|false|none|The short title of the incident|
|updated_at|string(date-time)|false|none|When the incident was last updated|

#### Enumerated Values

|Property|Value|
|---|---|
|Permissions|admin|
|Permissions|analyst|
|Permissions|guest|
|severity|0|
|severity|1|
|severity|2|
|severity|3|
|severity|4|
|severity|5|

<h2 id="tocS_NoMatch">NoMatch</h2>
<!-- backwards compatibility -->
<a id="schemanomatch"></a>
<a id="schema_NoMatch"></a>
<a id="tocSnomatch"></a>
<a id="tocsnomatch"></a>

```json
"Et dolor ipsam."

```

No vulnerability matched given criteria

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|No vulnerability matched given criteria|

<h2 id="tocS_SubmitRequestBody">SubmitRequestBody</h2>
<!-- backwards compatibility -->
<a id="schemasubmitrequestbody"></a>
<a id="schema_SubmitRequestBody"></a>
<a id="tocSsubmitrequestbody"></a>
<a id="tocssubmitrequestbody"></a>

```json
{
  "cvss_score": 0.8337925,
  "description": "Similique est dolor sit corporis excepturi quia.",
  "exploitable": false,
  "is_patchable": true,
  "is_upgradeable": false,
  "title": "Placeat modi debitis."
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|cvss_score|number(float)|false|none|Severity score of the vulnerability|
|description|string|false|none|Description of the vulnerability|
|exploitable|boolean|false|none|If the vulnerability is exploitable|
|is_patchable|boolean|false|none|If the vulnerability is patchable|
|is_upgradeable|boolean|false|none|If the vulnerability is upgradeable|
|title|string|false|none|Title of the vulnerability|

<h2 id="tocS_Vulnerability">Vulnerability</h2>
<!-- backwards compatibility -->
<a id="schemavulnerability"></a>
<a id="schema_Vulnerability"></a>
<a id="tocSvulnerability"></a>
<a id="tocsvulnerability"></a>

```json
{
  "cvss_score": 4.7,
  "description": "There is a possible out of bounds write due to a use after free. This could lead to remote code execution with no additional execution privileges needed.",
  "exploitable": true,
  "id": 1574086337224311300,
  "is_patchable": true,
  "is_upgradeable": false,
  "title": "CVE-2020-10"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|cvss_score|number(float)|false|none|Severity score of the vulnerability|
|description|string|false|none|Description of the vulnerability|
|exploitable|boolean|false|none|If the vulnerability is exploitable|
|id|integer|true|none|Unique ID of the vulnerability|
|is_patchable|boolean|false|none|If the vulnerability is patchable|
|is_upgradeable|boolean|false|none|If the vulnerability is upgradeable|
|title|string|false|none|Title of the vulnerability|

