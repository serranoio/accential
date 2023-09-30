import { LitElement, TemplateResult, html } from 'lit'
import { customElement, property, query } from 'lit/decorators.js'
import { AddMetricSteps, CreateMetricOptions, Metric, MetricSteps, dummyMetric, fromDocument, fromOutsideSource, setManually } from '../../../model/metric';
import createMetricCss from './create-metric.css';

import "../../shared/question.component"
import { SetSelectedTab } from '../../../model/events';
import { Distilled } from '../../../model/tabs';

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

  @property()
  chosenMethod = CreateMetricOptions.SetManually;

  @property()
  creatingMetricInputs: number = -1;

constructor() {
  super()

  }

  @property()
  addMetricSteps = AddMetricSteps.AddMetric;

  increaseMetricCounter() {
    this.metrics = [...this.metrics, structuredClone(dummyMetric)]

    this.dispatchEvent(new CustomEvent("update-metrics", {
      composed: true,
      bubbles: true,
      detail: {
        metrics: this.metrics
      }
    }))
  }
  
  decreaseMetricCounter() {
    if (this.metrics.length <= 1) {
      return;
    }
    
    this.metrics = this.metrics.slice(0, this.metrics.length-1)
    this.dispatchEvent(new CustomEvent("update-metrics", {
      composed: true,
      bubbles: true,
      detail: {
        metrics: this.metrics
      }
    }))
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
    this.metric.Label = e.target.querySelector("#Label").value
    this.metric.Explanation = e.target.querySelector("#Explanation").value
    
    this.metric.Rating = e.target.querySelector("#Rating").value


    this.dispatchEvent(new CustomEvent("add-new-metric", {
      composed: true,
      bubbles: true,
      detail: {
        metric: this.metric
      }
    }))

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
  
    this.dispatchEvent(new CustomEvent(SetSelectedTab, {
      bubbles: true,
      composed: true,
      detail: {
        selectedTab: Distilled
      }
    }))
  
    this.metricSteps = MetricSteps.EvaluateMetric;
    return "Evaluate Metric"
  }

  getDescription(name: string) {
    if (name === "Label") {
      return "This is the label"
    } else if (name === "Explanation") {
      return "Please explain the purpose or the significance of this metric"
    } else if (name === "Value") {
      return "Please input the value of the metric"
    } else if (name === "Rating") {
      return "Please give it your own rating. Optional."
    }
  }

  addParenthesisState(): TemplateResult {

    return html`
    <h2>Name Your Metric</h2>
    <p>Name your new metric, as well as add any necessary parenthesis for the calculation.</p>
    <div class="metric-div">
    <div class="labelinput">
    <label>Label
    <question-component
    .description=${this.getDescription("Label")}>
    .width=${20}
    </question-component>
    </label>
    <input type="text" id="Label" placeholder="Name of Metric">
    </div>
    <div class="labelinput">
    <label>Value
    <question-component
    .description=${this.getDescription("Value")}
    .width=${20}></question-component>
    </label>
    <input type="text" id="Value" value=${this.valueIntermediate}/>
    </div>
    <div class="labelinput">
    <label>Explanation
    <question-component
    .description=${this.getDescription("Explanation")}
    .width=${20}></question-component>
    </label>
    <input type="text" id="Explanation" placeholder="Explanation of Metric">
    </div>
    <div class="labelinput">
    <label>Rating
    <question-component
    .description=${this.getDescription("Rating")}
    .width=${30}></question-component>
  </label>
  <input type="number" id="Rating" placeholder="Rating">
</div>
</div>
`;
  }

  handleButtonClick(e: any): boolean {
    if (e.target.classList.contains("add-metric")) {
      return false;
    }

    return true;
  }

  sendNewStep(step: AddMetricSteps) {
    this.dispatchEvent(new CustomEvent("MetricSteps", {
      bubbles: true,
      composed: true,
      detail: {
        method: step
      }
    }))
  }

  setCurrentMetric(e: any) {
     this.creatingMetricInputs = e.target.closest("div").dataset.pos;

     this.dispatchEvent(new CustomEvent("creating-metric-inputs", {
      bubbles: true,
      composed: true,
      detail: {
        creatingMetricInputs: this.creatingMetricInputs
      }
    }))

  }

  switchAddMetricSteps(e: any) {
    
    if (this.addMetricSteps === AddMetricSteps.AddMetric) {
      this.setCurrentMetric(e)
      this.addMetricSteps = AddMetricSteps.ChooseMethod
      this.sendNewStep(AddMetricSteps.ChooseMethod)
    } else if (this.addMetricSteps === AddMetricSteps.ChooseMethod &&
                this.handleButtonClick(e)) {  // in HERE, we handle event

      this.sendNewStep(AddMetricSteps.AddingMetric)

      this.dispatchEvent(new CustomEvent(AddMetricSteps.AddingMetric, {
        bubbles: true,
        composed: true,
        detail: {
          method: e.target.innerHTML
        }
      }))


      this.addMetricSteps = AddMetricSteps.AddingMetric
    }
  }

  getAddMetricSteps(): TemplateResult {
    if (this.addMetricSteps === AddMetricSteps.AddMetric) {
      return html`Add Metric  <span> +<span>`
    } else if (this.addMetricSteps === AddMetricSteps.ChooseMethod) {   
      return html`
        <button type="button">${fromDocument}</button>
        <button type="button">${setManually}</button>
        <button type="button">${fromOutsideSource}</button>
      `
    } else {
      
      return html``
    }
  } 

  getAddMetricsButton(pos: number): TemplateResult {

    return html`<button
      type="button"
      style="cursor: ${this.addMetricSteps === AddMetricSteps.AddMetric ? "pointer" : ""};
           z-index: ${this.addMetricSteps === AddMetricSteps.AddingMetric && pos === Number(this.creatingMetricInputs) ? "-1" : "1"};
           box-shadow: ${this.addMetricSteps === AddMetricSteps.AddingMetric  && pos === Number(this.creatingMetricInputs) ? "0 0 0 1px var(--success)" : "0 0 0 0"};
           opacity: ${this.addMetricSteps === AddMetricSteps.AddingMetric  && pos === Number(this.creatingMetricInputs) ? "1" : ""};
      
           "
      @click=${this.switchAddMetricSteps}
      class="add-metric">${this.getAddMetricSteps()}
      </button>`
  }

  finishEditing() {
    this.addMetricSteps = AddMetricSteps.AddMetric
    this.sendNewStep(AddMetricSteps.AddMetric)
  }

  evaluateMetricsState(): TemplateResult[]  {
   return this.metrics.map((curMetric: Metric, pos: number) => {


      return html`
      ${pos === 0 ? "" : html`
      <div class="operation-div">
      <input type="text" placeholder="Operation"/>
      </div>
      `}
      <div class="metric-div" data-pos="${pos}">
        ${this.getAddMetricsButton(pos)}
      <div class="labelinput">
          <label>Label
            <question-component
            .description=${this.getDescription("Label")}
            .width=${20}
            ></question-component>
            </label>
            <input
            type="text"
            value=${curMetric.Label}
              placeholder="Label"/>
            </div>
            
            <div class="labelinput">
            <label>Value
            <question-component
            .description=${this.getDescription("Value")}
            .width=${20}
            ></question-component>
            </label>
            <input
            type="number"
            value=${curMetric.Value === -1 ? 0 : curMetric.Value}
              placeholder="Value"/>
            </div>
            
            <div class="labelinput">
            <label>Explanation
            <question-component
            .description=${this.getDescription("Explanation")}
            .width=${20}
            ></question-component>
            </label>
            <input
            
            value=${curMetric.Explanation}
             type="text"
              placeholder="Explanation"/>
            </div>
            ${this.addMetricSteps === AddMetricSteps.AddingMetric
              && pos === Number(this.creatingMetricInputs)
              ? html`
            <button class="finish-editing" type="button" @click=${this.finishEditing}>Finish >></button>
            </div>

            
            ` : ""}
            `
          });
          
        }
        
        render() {
 
          const addMetric = html` <button class="add-another-metric" @click=${this.increaseMetricCounter}> 
          Add Another Metric
          </button>`;
          
          const removeMetric = html` <button class="remove-another-metric" @click=${this.decreaseMetricCounter}> 
          Remove Another Metric
          </button>`
          
          const evaluateMetrics = this.evaluateMetricsState();
          
          const addParenthesisInput = this.addParenthesisState();
          
          const buttons = html`
          <div class="configure-metric-div">
          ${addMetric}
          ${removeMetric}
          </div>
          `


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
            <h2>Put together Submetrics</h2>
            <p>Put together multiple metrics from the document to create your own metric.</p>
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
