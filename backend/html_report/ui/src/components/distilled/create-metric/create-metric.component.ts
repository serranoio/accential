import { LitElement, html } from 'lit'
import { customElement, property, query } from 'lit/decorators.js'
import { Metric, MetricSteps, dummyMetric } from '../../../model/metric';
import createMetricCss from './create-metric.css';


// 3 Steps


@customElement('create-metric-component')
export class CreateMetricComponent extends LitElement {

  static styles = [createMetricCss]

  @property()
  metrics: Metric[] = [structuredClone(dummyMetric)]
  
  @property()
  metric: Metric = structuredClone(dummyMetric)

  @query('#metric-form')
  // @ts-ignore
  metricForm: HTMLFormElement;

  @property()
  valueIntermediate = "";
  
  @property()
  metricSteps: MetricSteps = MetricSteps.EvaluateMetric;

  constructor() {
    super()
}

  increaseMetricCounter() {
    this.metrics = [...this.metrics, structuredClone(dummyMetric)]
  }

  decreaseMetricCounter() {
    if (this.metrics.length <= 1) {
      return;
    }

    this.metrics = this.metrics.slice(0, this.metrics.length-1)
  }

  getMetric(position: number, metrics: Metric[]): string {
    if (position === metrics.length - 1) {
        return String(metrics[position].Value)
    }

    if (metrics[position].Operation === "+") {
      return metrics[position].Value + "+" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].Operation === "-") {
      return metrics[position].Value + "-" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].Operation === "x") {
      return metrics[position].Value + "*" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].Operation === "/") {
      return metrics[position].Value + "/" + this.getMetric(position+1, metrics);  
    } else {
      return String(metrics[position].Value);
    }


  }

  calculateMetricFromMetrics() {
    this.valueIntermediate = this.getMetric(0, this.metrics);
  }

  doTheMath(e: any) {
    this.metric.Value = eval(e.target.querySelector("#Value").value)

    console.log(this.metric.Value)

    this.metricSteps = MetricSteps.CreateMetric;
    // in this step, also include name of metric, explanation and rating
  }

  evaluateMetric() {
    let count = 0;
      
    const allInputs = this.metricForm.querySelectorAll("input");
    
    let curMetric = -1;
    for (const input of allInputs) {
      
      if (count % 4 === 0) {  // 0 mod 4 == 0
        curMetric++;
        this.metrics[curMetric].Label = input.value
      } else if (count % 4 === 1) {  // 1 mod 4 === 1
        this.metrics[curMetric].Value = Number(input.value)
      } else if (count % 4 === 2) {  // 2 mod 4 === 2
        this.metrics[curMetric].Explanation = input.value
      } else if (count % 4 === 3) {  // 3 mod 4 === 3
        if (curMetric < this.metrics.length - 1) {  // not the last one
          this.metrics[curMetric].Operation = input.value
        }
        
      }
      count++;  
    }
    this.metricSteps = MetricSteps.AddParenthesis;
    this.calculateMetricFromMetrics();
  }

  onCreateMetric(e: any) {
    e.preventDefault();


    if (this.metricSteps === MetricSteps.EvaluateMetric) {
        this.evaluateMetric()
    } else if (this.metricSteps === MetricSteps.AddParenthesis) {
      this.doTheMath(e);
    }

    // put metrics

  }

  buttonSteps(): string {

    if (this.metricSteps === MetricSteps.EvaluateMetric) {
      return "Evaluate Metric";
    } else if (this.metricSteps === MetricSteps.AddParenthesis) {
      return "Finalize Metric";
    }
  
    return "Done!";
  }

  render() {

    const addMetric = html` <button class="add-another-metric" @click=${this.increaseMetricCounter}> 
    Add Another Metric
    </button>`;

    const removeMetric = html` <button class="remove-another-metric" @click=${this.decreaseMetricCounter}> 
    Remove Another Metric
    </button>`
    
    const evaluateMetrics = this.metrics.map((_: Metric, pos: number) => {
            
      return html`
      ${pos === 0 ? "" : html`
      <div class="operation-div">
      <input type="text" placeholder="Operation"/>
      </div>
      `}
      <div class="metric-div">
      <input type="text" placeholder="Label"/>
      <input type="number" placeholder="Value"/>
      <input type="text" placeholder="Explanation"/>
      </div>
      `
    });

    const buttons = html`
    <div class="configure-metric-div">
    ${addMetric}
    ${removeMetric}
    </div>
    `

    const addParenthesisInput = html`
    <h2>Name Your Metric</h2>
    <p>Name your new metric, as well as add any necessary parenthesis for the calculation.</p>
    <div class="metric-div">
    <input type="text" placeholder="Name of Metric">
    <input type="text" id="Value" value=${this.valueIntermediate}/>
    <input type="text" placeholder="Explanation of Metric">
    <input type="number" placeholder="Rating">
    </div>
    `;

    return html`
        <figure
        class=""
        >
       ${buttons}
       
       <form
       @submit=${this.onCreateMetric}
       id="metric-form"
       >
       ${this.metricSteps === MetricSteps.AddParenthesis ? addParenthesisInput : ""}

          ${this.metricSteps === MetricSteps.EvaluateMetric ?
            html`
            <h2>Evaluate Your Metric</h2>
            <p>Put together multiple metrics from the document to create your own metric.<p>
            ${evaluateMetrics}
            `
             : 
            ""
          }

          <button
          type="submit"
          class="create-metric">
         ${this.buttonSteps()}</button>



        </form>

        </figure>
    `
  }
}
