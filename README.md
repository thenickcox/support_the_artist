## Support the Artist

A little command line tool that is meant to be run on a monthly crontab that, given your Last.fm
credentials, outputs the album you've listened to in the last month, and reminds you to
support the artist.

### Usage

Obtain a last.fm API key. Then run:

```
$ LASTFM_KEY=yourapikey LASTFM_USER=yourlastfmuser ./support_the_artist
# => Your most played album in the last month is "Have You In My Wilderness" by Julia Holter, which you've played 66 times.
     Buy the album and support the artist!
```

### Roadmap

* Configurable messages: choose to receive via email or SMS

A work in progress.
