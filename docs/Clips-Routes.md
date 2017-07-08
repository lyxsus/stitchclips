| Enpoints | Description |
| --- | --- |
| [Get clips](#get-clips) | Returns clips you asked for |


### Get clips

#### URL

`GET /clips/{channel}/{period}/{limit}`

#### Query Parameters

| Name | Type | Description |
| --- | --- | --- |
| `channel` | string | The Twitch channel you wish to get the  clips from
| `period` | string | The window of time to search for Clips. Valid values: `day`, `week`, `month`, `all`. 
| `limit` | integer | The number of clips you want to get

#### Example request

```bash
curl -X GET 'http://api.stitch.sadzeih.com/clips/itmejp/week/2'
```

#### Example response

```json
{
  "clips": [
    {
      "slug": "DaintyRenownedBubbleteaFutureMan",
      "tracking_id": "79284967",
      "url": "https://clips.twitch.tv/DaintyRenownedBubbleteaFutureMan?tt_medium=clips_api&tt_content=url",
      "embed_url": "https://clips.twitch.tv/embed?clip=DaintyRenownedBubbleteaFutureMan&tt_medium=clips_api&tt_content=embed",
      "embed_html": "<iframe src='https://clips.twitch.tv/embed?clip=DaintyRenownedBubbleteaFutureMan&tt_medium=clips_api&tt_content=embed' width='640' height='360' frameborder='0' scrolling='no' allowfullscreen='true'></iframe>",
      "broadcaster": {
        "id": "10817445",
        "name": "itmejp",
        "display_name": "itmeJP",
        "channel_url": "https://www.twitch.tv/itmejp",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/itmejp-profile_image-64703923f21827e3-150x150.png"
      },
      "curator": {
        "id": "66090647",
        "name": "noorelbahrain",
        "display_name": "NoorElBahrain",
        "channel_url": "https://www.twitch.tv/noorelbahrain",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/noorelbahrain-profile_image-13ca717bebfc13f2-150x150.jpeg"
      },
      "vod": {
        "id": "144686844",
        "url": "https://www.twitch.tv/videos/144686844?t=2h44m1s"
      },
      "game": "Injustice 2",
      "language": "en",
      "title": "Stop! Aquaman please stop!",
      "views": 812,
      "duration": 17.083984,
      "created_at": "2017-05-18T04:42:37Z",
      "thumbnails": {
        "medium": "https://clips-media-assets.twitch.tv/25303070832-offset-9841.750999999998-17.08333333333331-preview-480x272.jpg",
        "small": "https://clips-media-assets.twitch.tv/25303070832-offset-9841.750999999998-17.08333333333331-preview-260x147.jpg",
        "tiny": "https://clips-media-assets.twitch.tv/25303070832-offset-9841.750999999998-17.08333333333331-preview-86x45.jpg"
      }
    },
    {
      "slug": "RepleteSneakyTroutBudStar",
      "tracking_id": "80120379",
      "url": "https://clips.twitch.tv/RepleteSneakyTroutBudStar?tt_medium=clips_api&tt_content=url",
      "embed_url": "https://clips.twitch.tv/embed?clip=RepleteSneakyTroutBudStar&tt_medium=clips_api&tt_content=embed",
      "embed_html": "<iframe src='https://clips.twitch.tv/embed?clip=RepleteSneakyTroutBudStar&tt_medium=clips_api&tt_content=embed' width='640' height='360' frameborder='0' scrolling='no' allowfullscreen='true'></iframe>",
      "broadcaster": {
        "id": "10817445",
        "name": "itmejp",
        "display_name": "itmeJP",
        "channel_url": "https://www.twitch.tv/itmejp",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/itmejp-profile_image-64703923f21827e3-150x150.png"
      },
      "curator": {
        "id": "72020616",
        "name": "ravingsockmonkey",
        "display_name": "ravingsockmonkey",
        "channel_url": "https://www.twitch.tv/ravingsockmonkey",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/ravingsockmonkey-profile_image-0674d0ef4e885666-150x150.jpeg"
      },
      "vod": {
        "id": "145592193",
        "url": "https://www.twitch.tv/videos/145592193?t=2h21m21s"
      },
      "game": "Dungeons & Dragons",
      "language": "en",
      "title": "That's one way of dealing with things. @rollplay",
      "views": 57,
      "duration": 27.766992,
      "created_at": "2017-05-20T22:12:36Z",
      "thumbnails": {
        "medium": "https://clips-media-assets.twitch.tv/25322876720-offset-8481.240333333333-27.75-preview-480x272.jpg",
        "small": "https://clips-media-assets.twitch.tv/25322876720-offset-8481.240333333333-27.75-preview-260x147.jpg",
        "tiny": "https://clips-media-assets.twitch.tv/25322876720-offset-8481.240333333333-27.75-preview-86x45.jpg"
      }
    }
  ],
  "_cursor": "Mg=="
}
```