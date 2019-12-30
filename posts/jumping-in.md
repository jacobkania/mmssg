---
Title: Jumping In
PublishedDate: 2019-12-30
---
Writing the software to run a blog was a big challenge for me. Not because writing the software itself was challenging, but because I was paralyzed by choice of the best way to do it.

I had a whole list of requirements for the blog website. It had to

1. Support markdown for text input.
2. Allow me to add posts easily, and edit everything about them.
3. Be secure, without requiring any constant security updates or ongoing maintenance.
4. Just work. I didn't want to be constantly fiddling with settings.
5. Allow me to customize anything about it that I wanted. Colors, themes, etc etc etc.

I knew about *github pages* existing, but wasn't sure the best way to get a site hosted on there. Jekyll was the default that they offered, but seemed too weak to fit most of my needs without lots of customisation. I didn't want to learn to use static site generators that already existed -- they were out there, but it was hard for me to find simple examples that showcased how they worked. So I wrote my own "blog engine" (bla) in Winter of 2018.

It was fairly simple, but included a few good features: auth for admin users to write posts, a frontend ui for writing posts, and a way to host content on customisable HTML pages. The biggest downside though, was that this was a Go application. It needed to be running 24/7 on a server, and wasn't easily scalable. This would be a big problem for me, since I expected occasional spikes in traffic, while normally very minimal regular visitors. This would mean that I'd need to spend a lot of money on a powerful enough server for those unexpected spikes, while it would sit idle for days or weeks without much traffic at all.

I ended up abandoning Bla after that realization, and because I found a platform that fit many of my needs: [Svbtle](https://svbtle.com). This was a fully managed blog site, very minimalist themed (which I like), and allowed for markdown in the posts. It automatically scaled, and cost me $6.00 per month. Not too bad. But it didn't check all of my boxes, and I didn't like how they organized my index page. It was good enough for a while, but eventually I felt unmotivated to write there. It didn't feel like it was **my blog**.

Finally, I figured I'd make my own static site generator. It would be simple to use, easy to template for, and I could host my entire blog on github pages. This allowed me to hit all 5 of my requirements. I also hadn't written go in a while, so this was a good way to get back into it.