import { customElement, property } from 'lit/decorators.js'
import { Statistics, dummyStatistics } from '../../../model/statistics';
import { LitElement, html } from 'lit'



@customElement('metric-component')
export class MetricComponent extends LitElement {
@property()
statistics: Statistics = dummyStatistics;

@property()
openView = false;

@property()
label: string = "";

@property()
value: number = 0;

@property()
explanation: string = "";
    
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
    // <p>${this.Rating ? "Good" : "Bad"}</p>
    return html`
        <tr class="row second">
        <td class="column">
        <h3>${this.label}</h3>
        </td>
        <td class="column">
        <p>${this.value}</p>
        </td>
        <td class="column">
        none
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
        <p>${this.explanation}</p>
        </td>
        </tr>
    `
  }
}
