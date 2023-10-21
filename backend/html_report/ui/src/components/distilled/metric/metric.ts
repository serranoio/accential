import { customElement, property } from 'lit/decorators.js'
import { LitElement, html } from 'lit'
import { CreateMetricOptions, Metric, Submetric, dummyMetric } from '../../../model/metric';
import { DeleteMetric, EditMetric, UseMetric } from '../../../model/events';

@customElement('metric-component')
export class MetricComponent extends LitElement {

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

  this.dispatchEvent(new CustomEvent(UseMetric, {
      composed: true,
      bubbles: true,
      detail: {
          metric: this.metric
      }
    }))
  }
  
  editMetric() {
    this.dispatchEvent(new CustomEvent(EditMetric, {
      composed: true,
      bubbles: true,
      detail: {
          position: this.position,
          metric: this.metric
      }
    }))
  }

  deleteMetric() {
    this.dispatchEvent(new CustomEvent(DeleteMetric, {
      composed: true,
      bubbles: true,
      detail: {
          metric: this.metric
      }
    }))
  }
  
  render() {
    return html`
        <tr class="row second">
        <td class="column">
        <h3>${this.metric.label}</h3>
        </td>
        <td class="column">
        <p>${this.metric.value === -1 ? "could not find" : this.metric.value}</p>
        </td>
        <td class="column">
        n/a 
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
        <td>
        <p>${this.metric.explanation}</p>
        <p>Calculated using: </p>
        <ol>
        ${this.metric.submetric.map((submetric: Submetric) => {
          return html`<li><p>${submetric.label} with value of: ${submetric.value} ${submetric.operation}</p></li>`
        })}</ol>
        <button class="edit-metric" @click=${this.editMetric}>Edit</button>
        <button class="delete-metric" @click=${this.deleteMetric}>Delete</button>
        </td>
        </tr>
    `
  }
}
