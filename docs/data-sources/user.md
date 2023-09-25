---
page_title: "Data Source: okta_user"
description: |-
  Get a single users from Okta.
---

# Data Source: okta_user

Get a single users from Okta.

## Example Usage

```terraform
# Get a single user by their id value
data "okta_user" "example" {
  user_id = "00u22mtxlrJ8YkzXQ357"
}

# Search for a single user based on supported profile properties
data "okta_user" "example" {
  search {
    name  = "profile.firstName"
    value = "John"
  }

  search {
    name  = "profile.lastName"
    value = "Doe"
  }
}

# Search for a single user based on a raw search expression string
data "okta_user" "example" {
  search {
    expression = "profile.firstName eq \"John\""
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `compound_search_operator` (String) Search operator used when joining mulitple search clauses
- `delay_read_seconds` (String) Force delay of the user read by N seconds. Useful when eventual consistency of user information needs to be allowed for.
- `search` (Block Set) Filter to find user/users. Each filter will be concatenated with the compound search operator. Please be aware profile properties must match what is in Okta, which is likely camel case. Expression is a free form expression filter https://developer.okta.com/docs/reference/core-okta-api/#filter . The set name/value/comparison properties will be ignored if expression is present (see [below for nested schema](#nestedblock--search))
- `skip_groups` (Boolean) Do not populate user groups information (prevents additional API call)
- `skip_roles` (Boolean) Do not populate user roles information (prevents additional API call)
- `user_id` (String) Retrieve a single user based on their id

### Read-Only

- `admin_roles` (Set of String)
- `city` (String)
- `cost_center` (String)
- `country_code` (String)
- `custom_profile_attributes` (String)
- `department` (String)
- `display_name` (String)
- `division` (String)
- `email` (String)
- `employee_number` (String)
- `first_name` (String)
- `group_memberships` (Set of String)
- `honorific_prefix` (String)
- `honorific_suffix` (String)
- `id` (String) The ID of this resource.
- `last_name` (String)
- `locale` (String)
- `login` (String)
- `manager` (String)
- `manager_id` (String)
- `middle_name` (String)
- `mobile_phone` (String)
- `nick_name` (String)
- `organization` (String)
- `postal_address` (String)
- `preferred_language` (String)
- `primary_phone` (String)
- `profile_url` (String)
- `roles` (Set of String)
- `second_email` (String)
- `state` (String)
- `status` (String)
- `street_address` (String)
- `timezone` (String)
- `title` (String)
- `user_type` (String)
- `zip_code` (String)

<a id="nestedblock--search"></a>
### Nested Schema for `search`

Optional:

- `comparison` (String)
- `expression` (String) A raw search expression string. This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
- `name` (String) Property name to search for. This requires the search feature be on. Please see Okta documentation on their filter API for users. https://developer.okta.com/docs/api/resources/users#list-users-with-search
- `value` (String)

