import { LitElement, TemplateResult, html } from 'lit'
import { customElement, property, query } from 'lit/decorators.js'
import { AddMetricSteps, CreateMetricOptions, Metric, MetricSteps, Submetric, dummyMetric, dummySubmetric, fromDocument, fromOthers, fromOutsideSource, setManually } from '../../../model/metric';
import createMetricCss from './create-metric.css';

import "../../shared/question.component"
import { AddNewMetricEvent, ChangeMetricSteps, SetCurrentInput, SetSelectedTab, UpdateSubmetrics } from '../../../model/events';
import { Distilled } from '../../../model/tabs';
import { SetEq } from '../../../model/worker';

// 3 Steps


@customElement('create-metric-component')
export class CreateMetricComponent extends LitElement {

  static styles = [createMetricCss]

  @property()
  submetrics: Submetric[] = [structuredClone(dummySubmetric)]
  
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

  updateSubmetrics() {
    this.dispatchEvent(new CustomEvent(UpdateSubmetrics, {
      composed: true,
      bubbles: true,
      detail: {
        metrics: this.submetrics
      }
    }))
  }

  increaseMetricCounter() {
    this.submetrics = [...this.submetrics, structuredClone(dummySubmetric)]

    this.updateSubmetrics();
  }
  
  decreaseMetricCounter() {
    if (this.submetrics.length <= 1) {
      return;
    }
    
    this.submetrics = this.submetrics.slice(0, this.submetrics.length-1)
    this.updateSubmetrics();
  }

  getMetric(position: number, metrics: Submetric[]): string {
    if (position === metrics.length - 1) {
        return String(metrics[position].value)
    }

    if (metrics[position].operation === "+") {
      return metrics[position].value + "+" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].operation === "-") {
      return metrics[position].value + "-" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].operation === "*") {
      return metrics[position].value + "*" + this.getMetric(position+1, metrics);  
    } else if (metrics[position].operation === "/") {
      return metrics[position].value + "/" + this.getMetric(position+1, metrics);  
    } else {
      return String(metrics[position].value);
    }
  }

  calculateMetricFromMetrics() {
    this.valueIntermediate = this.getMetric(0, this.submetrics);
  }

  doTheMath(e: any) {
    const eq = e.target.querySelector("#Value").value
    this.metric.value = eval(eq)
    this.metric.label = e.target.querySelector("#Label").value
    this.metric.explanation = e.target.querySelector("#Explanation").value
    this.metric.rating = e.target.querySelector("#Rating").value
    this.metric.submetric = this.submetrics
    this.metric.operation = SetEq(eq)

    this.dispatchEvent(new CustomEvent(AddNewMetricEvent, {
      composed: true,
      bubbles: true,
      detail: {
        metric: this.metric,
      }
    }))

    this.metricSteps = MetricSteps.CreateMetric;
  }

  getInputs() {
    let count = 0;
      
    const allInputs = this.metricForm.querySelectorAll("input");
    
    let curMetric = -1;
    for (const input of allInputs) {
      
      if (count % 4 === 0) {  // 0 mod 4 == 0
        curMetric++;
        this.submetrics[curMetric].order = curMetric
        this.submetrics[curMetric].label = input.value
      } else if (count % 4 === 1) {  // 1 mod 4 === 1
        this.submetrics[curMetric].value = Number(input.value)
      } else if (count % 4 === 2) {  // 2 mod 4 === 2
        this.submetrics[curMetric].explanation = input.value
      } else if (count % 4 === 3) {  // 3 mod 4 === 3
        if (curMetric < this.submetrics.length - 1) {  // not the last one
          this.submetrics[curMetric].operation = input.value
        }
        
      }
      count++;  
    }

    this.updateSubmetrics()
  }

  evaluateMetric() {
    this.getInputs()
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

  handleButtonClick(e: any): boolean {
    if (e.target.classList.contains("add-metric")) {
      return false;
    }

    return true;
  }

  sendNewStep(step: AddMetricSteps) {
    this.dispatchEvent(new CustomEvent(ChangeMetricSteps, {
      bubbles: true,
      composed: true,
      detail: {
        method: step
      }
    }))
  }

  setCurrentMetric(e: any) {
     this.creatingMetricInputs = e.target.closest("div").dataset.pos;

     this.dispatchEvent(new CustomEvent(SetCurrentInput, {
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
      return html`Edit Metric  <span> +<span>`
    } else if (this.addMetricSteps === AddMetricSteps.ChooseMethod) {   
      return html`
        <button type="button">${fromDocument}</button>
        <button type="button">${setManually}</button>
        <button type="button">${fromOutsideSource}</button>
        <button type="button">${fromOthers}</button>
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
    this.getInputs()
  }
  
  onOperationChange(e: any) {
    this.submetrics[e.target.closest("div").dataset.pos-1].operation = e.target.value
    this.getInputs()
  }
  

  evaluateMetricsState(): TemplateResult[]  {
   return this.submetrics.map((curMetric: Submetric, pos: number) => {

      return html`
      ${pos === 0 ? "" : html`
      <div class="operation-div" data-pos="${pos}">
      <input type="text" 
      placeholder="Operation"
      value=${this.submetrics[pos-1].operation}
      @input=${this.onOperationChange}
      />

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
            value=${curMetric.label}    
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
            value=${curMetric.value === -1 ? 0 : curMetric.value}
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
            
            value=${curMetric.explanation}
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

  addParenthesisState(): TemplateResult {

    return html`
    <h2>Name Your Metric</h2>
    <p>Name your new metric, as well as add any necessary parenthesis for the calculation.</p>
    <div class="metric-div">
    <div class="labelinput">
    <label>Label
    <question-component
    .description=${this.getDescription("Label")}
    .width=${15}
    >
    </question-component>
    </label>
    <input type="text" id="Label" value=${this.metric.label} placeholder="Name of Metric">
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
    <input type="text" id="Explanation" value=${this.metric.explanation} placeholder="Explanation of Metric">
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
            ${this.buttonSteps()}
          </button>
        </form>
        </figure>
    `
  }
}
