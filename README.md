# envidict

An English-Vietnamese dictionary website

## How to run?

```console
docker build -t envidict .
docker run -it --rm -p 3000:3000 -p 5000:5000 --name envidict envidict
```

Access the app at [http://localhost:3000](http://localhost:3000).
