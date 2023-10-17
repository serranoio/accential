to-do:

1. Server built lit element content from backend into an iframe on the frontend. Have it be VERY interactive.
    2. this will be a full stack single page framework, where I can communicate between different iFrames
        3. https://stackoverflow.com/questions/9153445/how-to-communicate-between-iframe-and-the-parent-site


4. No need to look at how to hydrate document yet.
5. First styling.
6. Change the frontend to orange color
7. In the Let's party section, make it a warehouse color for document insertion.

8. Make document look sexy with full reactivity and pagefind
    a. Metric: Title, Metric No., Rating, More,
    b. In "More" section,the keywords are buttons. you can hover over them to see the pagefind
    c. report



Sept 23rd.
Gym,
shower,
lunch,
call parents.

Investing research!!!
Dinner

Sunday:
Spanish Homework
Do coding.

Monday:
Study

Tuesday:
None

Wednesday:
Study

### LATER
 create notification channel

In rust:
 2. Add 3 steps:
    1. Choose a Company
    2. Choose metrics
    3. Send
- 1. Choose document
    - show available documents
        - create get all documents endpoint
    - or upload manually
- 2. Choose metrics
- 3. Send :}
2. Update post request to include metrics
 Use serde to serialize a struct from rust into json, include metrics and 
 file

 3. Make sure your document is saved

 4. Now we can make tablizer better.
 5. include <font> in 


### 1

POSS... 
in "Choose Saved"
adding metric will make post request instead of get request to 
add metrics ???
problem is that you will have to use financer again on tablizer


get document

get reportBytes...
process it using tablizer

Financizer with THOSE metrics to add.

theres no need to call create_report

just populate metrics with doc


document.metrics = append(document.metrics, InitFinancer()); 

save

send back

### 2

The add metric route needs to include the document ID to be added to the document ID

add order to save submetrics



// breakpoint at line 103

// when you find td, you consume next text token which is the contents

// don't consume next content


arelleCmdLine --plugins=/Users/davidserrano/Library/Python/3.9/lib/python/site-packages/iXBRLViewerPlugin -f /Users/davidserrano/greatness/rust/accential/backend/10-k-example.html --save-viewer ixbrl-report.html --viewer-url /Users/davidserrano/greatness/rust/accential/backend/html_report/ui/public/xbrl/ixbrlviewer.js


# Deploy

* Deployed onto fly using flyctl
* Backend deployed at https://accential.fly.dev
* used fly launch to launch new app
* used fly deploy to update docker image

# Steps
* Change url in api.ts to accential.dev