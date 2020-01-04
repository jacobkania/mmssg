# MMSSG
Most Minimal Static Site Generator

Hi there! I'm Jacob. Glad you stopped by the repo for MMSSG. This whole project started because I wasn't happy with any of the solutions out there for hosting my own blog. It had to be dead simple, basically just plug-and-play with a folder of markdown files with a template to plug the content into.

# Using MMSSG
Prerequisites:
* Some understanding of git -- making your own project, etc. -- and also git must be installed locally [(see here)](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* A version of Go installed on your computer to execute the site generator [(see here)](https://golang.org/doc/install)
* About 30 minutes to customize things and make them yours

If you would like to use this project for making your own website, congrats, it's already most of the way there! 

Start by forking this project on github, then clone your new repo to your computer.

Blog posts are written in markdown, and stored in the 'posts' folder in my simple configuration. Every file in this folder will become a page on your blog. The file names become the URL for each page, and each file has a header that contains metadata used while rendering the pages. Go ahead and delete all but one, and replace it with your own first blog post! Don't forget to rename the file too.

The other big customization that you'll want to do is editing the templates that your site uses. In the root folder, there are two files: `index.html` and `page.html`. These correspond to the two types of pages that we'll need templates for: the homepage with a list of all your posts, and the pages that contain the posts themselves.

The `index.html` template will only be used once: for your blog's landing page. However, the `page.html` will be used for each post that you write, containing its content. Anything wrapped in `{{}}` is where the templating part comes into play -- rendering different content depending on your posts. There are just a few key template tags that you'll need to stick into your page template, and those are all used in the default templates.

Go ahead and delete the 'docs' folder at the root of this project -- it contains the output of the demo site. This folder will be created again when we run the generator at the end. MMSSG is non-desctructive, so it won't clear out the dead pages in here that you won't need for your site.

After writing your own post, and editing the html templates to whatever you want, run this command from the root of the project: `go run main.go -i posts -u mmssg`. This will generate the blog from your posts and templates, and output to the 'docs' folder. The `-u mmssg` is for the leading URL of your site: ie the name of your repo. If you named your repo something different, you should specify the name of your repo in the `-u` flag. (This is because of how github pages creates the URL for your site).

Using Git, push up your changes to your github repo. Go to the 'Settings' tab within your repo, and scroll down to the section titled 'GitHub Pages'. In the dropdown for 'Source', choose 'master branch /docs folder'. This publishes your repo on github pages using the content in the /docs folder.

That's it! Your site is published.
