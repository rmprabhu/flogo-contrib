# tibco-twilio
This activity provides your flogo application the ability to call via Twilio.


## Installation

```bash
flogo add activity github.com/rmprabhu/flogo-contrib/activity/twilio-call
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "accountSID",
      "type": "string"
    },
    {
      "name": "authToken",
      "type": "string"
    },
    {
      "name": "from",
      "type": "string"
    },
    {
      "name": "to",
      "type": "string"
    },
    {
      "name": "url",
      "type": "string"
    }  
  ],
  "outputs": []
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| accountSID | The Twilio account SID |         
| authToken  | The Twilio auth token  |
| from       | The Twilio number used to originate the call from |
| to         | The number you are calling to |
| message    | The callback url with call parameters in twiml eg - https://demo.twilio.com/docs/voice.xml |
Note:
Phone numbers should be in the format '+15555555555'

## Configuration Examples
### Simple
Configure a task in flow to send 'my text message' to '617-555-5555' via Twilio:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-twilio-call",
  "name": "Call using Twillio",
  "attributes": [
    { "name": "accountSID", "value": "A...9" },
    { "name": "authToken", "value": "A...9" },
    { "name": "from", "value": "+12016901385" },
    { "name": "to", "value": "+16175555555" },
    { "name": "url", "value": "https://demo.twilio.com/docs/voice.xml" }
  ]
}
```

### Advanced
Configure a task in flow to send 'my text message' to a number from a REST trigger's query parameter:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-twilio",
  "name": "Send Text Message",
  "attributes": [
    { "name": "accountSID", "value": "A...9" },
    { "name": "authToken", "value": "A...9" },
    { "name": "from", "value": "+12016901385" },
    { "name": "url", "value": "https://demo.twilio.com/docs/voice.xml" }
  ],
  "inputMappings": [
    { "type": 1, "value": "[T.queryParams].From", "mapTo": "to" }
  ]
}
```
