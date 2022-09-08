# Wiki API

Get short descriptions for any Wikipedia page

## API docs

http://localhost:9090/docs

## Examples

Default lang = en

http://localhost:9090/page/Berlin


```
{
"title": "Berlin",
"lang": "en",
"description": "Capital and largest city of Germany",
"source": "local"
}
```

Get description, with fallback to Wikidata if there is no SHORTDESC in the Wikipedia page:

http://localhost:9090/page/Berlin?lang=de&fallback=true


```
{
"title": "Berlin",
"lang": "de",
"description": "Hauptstadt und Land der Bundesrepublik Deutschland",
"source": "central"
}
```