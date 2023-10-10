# October 9th, 2023.
- MVP is NOT done.
- Creating bare metrics
- fixed hover animation

# October 8th, 2023.
## Status 
- The MVP is *done*. No more features to be added
## TO:DO
- We just need to write more sets of metrics
    - Evaluation
    - Bare Metrics
    - Status
- gather folder of 10-k documents for testing
    ### WRITE TESTS FOR DOCUMENTS. MAKE SURE TABLE HAS
    ### all information
    - test on other documents
    - balance sheets
- Hydrate font into document
- Make video
    - you already wrote the script on your phone

# October 7th, 2023.
- finished 3 steps in Rust frontend ;)
- now working on making the parser work for all documents

# October 4th, 2023.
- removed rabbitmq. Now you can hit all api endpoints in success
- added comm package for types (communication)
- Added change name endpoint
- Starting to create Logic in frontend to parse/choose files


# October 3th, 2023.
- Added cool font. Also made world SVG move.
- Thought I should have a changelog.


## feature list:

- Create your own metrics
- add bare metrics from frontned
- add capability to pull documents straight from
    searching on sec.gov. So, we will embed sec.gov
    into the website and it will be used as a variant
    of Step 1.
- Multiple selection in doucment
- DEPLOY
- EDIT METRICS
- Use other metrics for calculation

 ChooseMethod options, 
    add method "FromOthers".
    Move over to metrics tab
    Include (+) button on each top level metric
    Pressing it will take Label + Value + Explanation from metric
    and input it into input field.

    Question????
    
    The metric will be based on OTHER metrics. It will be LINKED to the other metrics

    If metric contains link, check links to find other metric value and use calculation for that
    instead of withi the tables

### You can also edit old metrics.
# Edit button shows up under "more"

## bugs
Still cannot choose a metric from an earlier year
