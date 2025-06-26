# Go Web Server - My Learnings

My goal was to get started with go with this particular project.

What I learnt while building this:  
- Go's basic syntax, structs, functions

- net/HTTP and how do we handle the handlers, I liked doing the routing part

- Mux I did learn about mux.

  - I was actually curious about why was there a nil on the side of port in ListenAndServe() so searched on it and got to learn about multiplexers which was amazing to know. I did an iteration where I routed the whole thing with mux1 and mux2 and a main mux serving them, it was great learning but i got out of it.

The most difficult part in this particular project was the contact form.

The initial plan of just showing the form details on the contact page using Fprintf was very basic to me, so I went on to build it inside the html page itself. Honestly it felt like I shouldnt have picked it up but yeah, I did and learnt about go Templating.

I am able to understand the flow of it but I have the hunch that I might forget it. This was to put the text on the html page. 3-4 things in that were:

- ParseForms
- The Execute method
- Populating the data struct which had the details

One dumb thing I was stuck on but still was necessary was the attributes in the html form. The name="name" was taken by me in the contactHandler data assigning part as Name, which made the issue coz the name value pair which Form in go uses couldnt find the same attribute.

## Other Notes

Apart from that I would not take the credit of whatever css is done on it. After I did the whole project, I went on each page and made gemini stylize the pages so that atleast I come back on the pages lol.

Even though this was just a warmup as suggested by gemini to learn to build something for the first time in golang, I built more than that was asked, and I still wish that there was a way I could store the form details and manage them, so maybe I will do a CRUD app next.
