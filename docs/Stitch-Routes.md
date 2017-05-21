| Enpoints | Description |
| --- | --- |
| [Stitch](#stitch) | Returns the URL to the stitched video |


### Stitch

#### URL

`POST /stitch`

#### Request Body

In the request body you must provide the slugs for the clips you want to stitch together.

The request body must be in JSON.

##### Example body

```json
{
	"clips": [
		{
			"slug": "DaintyRenownedBubbleteaFutureMan"
		},
		{
			"slug": "TriangularCleanSardineOMGScoots"
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
-d '{"clips":[{"slug":"DaintyRenownedBubbleteaFutureMan"},{"slug":"TriangularCleanSardineOMGScoots"},{"slug":"RepleteSneakyTroutBudStar"}]}' \
-X POST 'https://apidomain.com/stitch'
```


#### Example response

```json
{
  "id": "cf4f2f50-3e3e-11e7-b8eb-1c872c71269a",
  "url": "http://localhost:8000/video/cf4f2f50-3e3e-11e7-b8eb-1c872c71269a.mp4"
}
```