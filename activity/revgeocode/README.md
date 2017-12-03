# tibco-revgeocode
This activity provides your flogo application the ability to get address from coordinates using google maps API.


## Installation

```bash
flogo add activity github.com/rprabhu/flogo-contrib/activity/revgeocode
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "apiKey",
      "type": "string"
    },
    {
      "name": "lat",
      "type": "string"
    },
	{
      "name": "long",
      "type": "string"
    }  
  ],
  "outputs": [
     {
       "name": "location",
       "type": "string"
     }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| apiKey      | The Google Maps API key |         
| lat         | Latitude  |
| long        | Longitude |

## Configuration Examples
### Simple
Configure a task in flow to get address from co-oordinates via Google Maps API:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-revgeocode",
  "name": "Get Location",
  "attributes": [
    { "name": "apiKey", "value": "A...9" },
    { "name": "lat", "value": "40.33...." },
    { "name": "long", "value": "-10.99....." }
  ]
}
```

### Advanced
Configure a task in flow to get address from co-oordintes via Google Maps API from a REST trigger's query parameter:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-revgeocode",
  "name": "Get Location",
  "attributes": [
    { "name": "apiKey", "value": "A...9" }
  ],
  "inputMappings": [
    { "type": 1, "value": "[T.queryParams].lat", "mapTo": "lat" }
    { "type": 1, "value": "[T.queryParams].long", "mapTo": "long" }
  ]
}
```
