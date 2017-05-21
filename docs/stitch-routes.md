| Enpoints | Description |
| --- | --- |
| [Stitch](#stitch) | Returns the URL to the stitched video |


### Stitch

#### URL

`POST /stitch`

#### Request Body

The request body must be in JSON.
You must provide Clips slugs in this format:

```json
{
	clips: [
		{
			"slug": "DaintyRenownedBubbleteaFutureMan"
		},
		{
			"slug": "RepleteSneakyTroutBudStar"
		}
	]
}
```

#### Example request

```bash
curl -H 'Content-Type: application/json' \
-d '{clips: [{"slug": "DaintyRenownedBubbleteaFutureMan"},{"slug": "RepleteSneakyTroutBudStar"}]}' \
-X POST 'http://apidomain.com/stitch'
```


#### Example response

```json
{
	"URL": ""
}
```