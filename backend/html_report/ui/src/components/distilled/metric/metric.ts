import { customElement, property } from 'lit/decorators.js'
import { LitElement, TemplateResult, html } from 'lit'
import { CreateMetricOptions, Metric, Submetric, dummyMetric } from '../../../model/metric';
import { DeleteMetric, EditMetric, UseMetric } from '../../../model/events';
import { styleMap } from 'lit/directives/style-map.js';
import metricCss from './metric.css';

@customElement('metric-component')
export class MetricComponent extends LitElement {

  static styles = metricCss

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

  getStyles(rating: number): {backgroundColor: string, transform: string} {
    if (rating === 10) {
    return {
      backgroundColor: "var(--successColorD20);",
      transform: `scaleX(${rating / 10})`
    }
  } else if (rating === 9) {
    return {
      backgroundColor: "var(--successD10);",
      transform: `scaleX(${rating / 10})`
    }
  } else if (rating === 8) {
    return {
      backgroundColor: "var(--success);",
      transform: `scaleX(${rating / 10})`
    }
  } else if (rating === 7) {
      return {
        backgroundColor: "var(--warningColorD20);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 6) {
      return {
        backgroundColor: "var(--warningColorD10);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 5) {
      return {
        backgroundColor: "var(--warningColor);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 4) {
      return {
        backgroundColor: "var(--alertColorD10);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 3) {
      return {
        backgroundColor: "var(--alertColor);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 2) {
      return {
        backgroundColor: "var(--errorColor);",
        transform: `scaleX(${rating / 10})`
      }
    } else if (rating === 1) {
      return {
        backgroundColor: "var(--errorColorD10);",
        transform: `scaleX(${rating / 10})`
      }
    }
    return {
      backgroundColor: "",
      transform: "",
    }
  }

  createSVG(rating: string): TemplateResult {
    const styles = this.getStyles(Number(rating))

    return html`<div
    class="rating"
    style=${styleMap(styles)}
    >
    ${styles.transform === "" ? "N/A" : ""}
    </div>`
  }

  render() {

    return html`
        <tr class="row second">
        <td class="column">
        <h3>${this.metric.label}</h3>
        </td>
        <td class="column">
        ${this.createSVG(this.metric.rating)}
        </td>
        <td class="column">
        <p>${this.metric.value === -1 ? "could not find" : this.metric.value}</p>
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
