I'm looking to get into [Summits on the Air](https://www.sotadata.org.uk:443/summitslist.csv), a ham radio activity where you go to the top of a mountain and try to make contacts with people. Alternatively, you can be a "chaser" where you're looking for people who are on the mountains. You can also do both at the same time.

There's a database of all of the valid summits and everything available at [sotadata.org.uk](https://www.sotadata.org.uk:443/summitslist.csv). This is a big CSV file, however, and not really importable into mapping software. The [SOTA Mapping Project](https://www.sotamaps.org) has GPX and KML exports, but the GPX export doesn't have all of the summit stats (value, number of activations, etc) and the KML export has a whole bunch of stuff but formatted as HTML in the description field of the Placemark, which [Guru Maps](https://gurumaps.app), my preferred mapping software for things like this, doesn't handle well.

Rather than try to write some XSLT or some other thing to transform the KML into something that might work for me, I figure I can generate it from scratch manually if I have all the data from the CSV file. And maybe I can even make it templateable or something. I dunno.
