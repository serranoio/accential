import { LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import { Statistics, dummyStatistics } from '../../../model/statistics';



@customElement('metric-component')
export class MetricComponent extends LitElement {
@property()
statistics: Statistics = dummyStatistics;

@property()
openView = false;
    
  constructor() {
    super()
  }

  protected createRenderRoot(): Element | ShadowRoot {
      return this;
  }


  handleClick() {
    this.openView = !this.openView;
  }

  render() {
    return html`
        <tr class="row second">
        <td class="column">
        <h3>Working Capital</h3>
        </td>
        <td class="column">
        <p>${this.statistics.Mowc.WorkingCapital[0]}</p>
        </td>
        <td class="column">
          <p>${this.statistics.Mowc.WorkingCapital[0] > 2 ? "Good" : "Bad"}</p>
        </td>
        <td class="column">
        <button
        class="more ${this.openView ? "turn" : ""}"
        @click=${this.handleClick}
        >v</button>
        </td>
        </tr>
        
        <tr class="row info ${this.openView ? "show" : ""}">
        <td >
        <p>${this.statistics.Mowc.Info}<p>
        <p>Total Assets: $${this.statistics.Mowc.TotalAssets[0]}</p>
        <p>Total Liabilities: $${this.statistics.Mowc.TotalLiabilities[0]}</p>
        </td>
        </tr>
    `
  }
}
