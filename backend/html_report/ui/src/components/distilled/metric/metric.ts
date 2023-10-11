import { customElement, property } from 'lit/decorators.js'
import { Statistics, dummyStatistics } from '../../../model/statistics';
import { LitElement, html } from 'lit'
import { CreateMetricOptions, Metric, dummyMetric } from '../../../model/metric';



@customElement('metric-component')
export class MetricComponent extends LitElement {
@property()
statistics: Statistics = dummyStatistics;

@property()
openView = false;

@property()
position: number = 0;

@property()
metric: Metric = dummyMetric

@property()
chosenMethod = CreateMetricOptions.SetManually
    
  constructor() {
    super()
  }

  protected createRenderRoot(): Element | ShadowRoot {
      return this;
  }


  handleClick() {
    this.openView = !this.openView;
  }

  addMetric() {

    this.dispatchEvent(new CustomEvent("UseMetric", {
      composed: true,
      bubbles: true,
      detail: {
          metric: this.metric
      }
    }))
  }
  
  editMetric() {
    this.dispatchEvent(new CustomEvent("EditMetric", {
      composed: true,
      bubbles: true,
      detail: {
          position: this.position,
          metric: this.metric
      }
    }))
  }

  render() {
    // <p>${this.Rating ? "Good" : "Bad"}</p>
    return html`
        <tr class="row second">
        <td class="column">
        <h3>${this.metric.label}</h3>
        </td>
        <td class="column">
        <p>${this.metric.value}</p>
        </td>
        <td class="column">
        none
        </td>
        <td class="column">
        ${this.chosenMethod === CreateMetricOptions.FromOthers ?
          html`<button class="include-metric" @click=${this.addMetric}>+</button>` : ""}
        <button
        class="more ${this.openView ? "turn" : ""}"
        @click=${this.handleClick}
        >v</button>
        </td>
        </tr>
        
        <tr class="row info ${this.openView ? "show" : ""}">
        <td >
        <p>${this.metric.explanation}</p>
        <button class="edit-metric" @click=${this.editMetric}>Edit</button>
        </td>
        </tr>
    `
  }
}
