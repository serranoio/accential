import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import distilledCss from './distilled.css';
import "./metric/metric"
import "./create-metric/create-metric.component"
import { CreateMetricOptions, Metric } from '../../model/metric';
import { styleMap } from 'lit/directives/style-map.js';

@customElement('distilled-component')
export class DistilledComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [distilledCss] 

  @property()
  metrics: Metric[] = []

  @property()
  chosenMethod = CreateMetricOptions.SetManually

  constructor() {
    super()
  }

  calcRating() {
    if (!this.metrics) return {
      color: "",
      text: "",
    };

    let total = 0;
    let count = 0;
    for (let i = 0; i < this.metrics.length; i++) {
      if (this.metrics[i].rating !== "") {
        total += Number(this.metrics[i].rating)
        count++;
      }
    }

    const rating = total / count;

    if (rating >= 8) {
      return {
        background: `linear-gradient(
          90deg,
          var(--success),
          var(--successD10),
          var(--success)
        ) 0 0 / var(--bg-size) 100%;`,
        text: "Invest"
      }
    } else if (rating >= 6) {
      return {
        background: `linear-gradient(
          90deg,
          var(--warningColor),
          var(--warningColorD10),
          var(--warningColor)
        ) 0 0 / var(--bg-size) 100%;`,
        text: "Possible invest"}
    } else if (rating >= 3) {
      return {
        background: `linear-gradient(
          90deg,
          var(--alertColor),
          var(--alertColorD10),
          var(--alertColor)
        ) 0 0 / var(--bg-size) 100%;`,
        text: "Risk taker=invest"
      }
    } else if (rating >= 1) {
      return {
        background: `linear-gradient(
          90deg,
          var(--errorColor),
          var(--errorColorD10),
          var(--errorColor)
        ) 0 0 / var(--bg-size) 100%;`,
        text: "Don't Invest"
      }
    } else {
      return {
        background: "",
        text: "",
      }
    }

  }

  // if there are no votes, we need to say no rating yet
  displayRating() {
    const results = this.calcRating()

    console.log(results, this.metrics)

    return html`
    <p
    class="overall-rating"
    style=${styleMap({background: results.background, display: results.background === "" ? "none" : "block"})}
    >
    ${results.text}
    </p>
    `

  }

  render() {

    return html`
    <section class="full">
    <table class="report">
    <tbody>
    <tr class="row titles">
    <td class="column">
    <h3>Metric</h3>
    </td>
    <td class="column">
    <p>Rating</p>
    ${this.displayRating()}
       </td>
       <td class="column">
                <p>Value</p>
              </td>
              
              <td class="column">
              <p>More</p>
              </td>
            </tr>
          </tbody> 
        </table> 
        ${this.metrics.map((metric: Metric, position: number) => { 
          return html`
          <metric-component
          .metric=${metric}
          .position=${position}
          .chosenMethod=${this.chosenMethod}
          >
          </metric-component>
          `
        })}
       </section>
    `
  }
}



