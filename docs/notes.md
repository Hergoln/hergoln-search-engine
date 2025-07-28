Processes to cover based on [this articles][1]:
1. Crawling **the web**
  1. Be smart about it:
    * 500 HTTP errors mean service is overloaded
    * Run JavaScript as often pages use it to bring in majority of the content
    * Authorization and bots prevention issues
2. Indexing crawled web
  1. Classify to duplicate or canonical page
  2. Save information about page to database
    * Country of origin
    * Language of the page
    * Usability?
  3. Be mindfull of robots meta rules and other mechanisms that prevent bots from crawling to often
3. Serving search results based on indexed database
  1. Utilize users information to determine the search
  2. User's query matter: "bicycle shop" should return different results then "modern bicycles"

Articles:
* [How google search works][1]

[1]: <https://developers.google.com/search/docs/fundamentals/how-search-works> "How google search works article by google"
