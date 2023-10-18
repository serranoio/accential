import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import distilledCss from './distilled.css';
import "./metric/metric"
import "./create-metric/create-metric.component"
import { CreateMetricOptions, Metric } from '../../model/metric';

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
                <p>$</p>
              </td>
              <td class="column">
                <p>Rating</p>
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



