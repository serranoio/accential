#TODO

In the "$" column make it reactive:
    $ or Percentage or whatever

In the "Rating" column make it reactive:
    Rating or Valuation;
        For Valuation:
            Overvalued
            Undervalued
            Valued well


TAKE THE SPLUNK COLORS, add some brown to them lol.

The metric is calculated in UI when inputted manually. It is sent to DB
Financizer takes these metrics, finds them and calculates them.

Metric {
    Label:
    Value:
    Explanation:
    Operation:
    NextMetric:
}


PE Ratio * Earnings = Evaluation

Dividend Discount Model

Evaluation Multiple:

EV/Renuve
EV/EBITDA
P/E

Perform Fundamental Analysis of the company

- Include ticker symbol under the company name in the tab
- Retrieve stock price


- firstly, youre only able to click the row.

- when clicking label, you have options:
    - search for it in the document,
        - on this page, it shows you the metric youre creating. it moves from label,
        then to value, then explanation. Once finished, it takes you back to the create 
        metric screen and you see what you made.
    - set manually
    - retreive from outside source
        - stock price

Im gonna have to create global state
Do not create man purse.


State:
Creating a Metric
All Metrics

Retrieve from: 
IndexxedDB
    - Creating a Metric

Pocketbase
    - All Metrics
    - On start: grabAll
    - On every open: grabAll


Pocketbase records:
    - account
    - document name
    - metrics
    - document
    - name of take


Add tags to the metric name:
    Includes metric from doc
    includes metric from doc




2 ways:

Try to feed metrics regularly

make dummy components

Now:

Made dummy component

Do:

<div class="metrics">  // metrics
    <figure>  // metric
        <p>
        <p>
        <p>
        <p>
    </figure>
    <figure>  // metric
        <p>
        <p>
        <p>
        <p>
    </figure>
</div>

<figure class="doc">

</figure>

DIVISION BY 0 ERROR

COMPARE EVERYTHING TO LOWER CASE




To choose default metrics, when creating a document in post route, we will have default metrics (nothing)
hitting the change metric route will update the value within  the document OR

ONE ROUTE, and we send multiple value ;)