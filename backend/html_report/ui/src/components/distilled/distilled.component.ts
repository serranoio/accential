import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import distilledCss from './distilled.css';
import { Statistics, dummyStatistics } from '../../model/statistics';

import "./metric/metric"
import "./create-metric/create-metric.component"
import { Metric, dummyMetric } from '../../model/metric';
import { AddNewMetric } from '../../model/api';
import { getDocId } from '../../model/util';

@customElement('distilled-component')
export class DistilledComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [distilledCss] 

  @property()
  statistics: Statistics = dummyStatistics;

  @property()
  metrics: Metric[] = []
  
  @property()
  newMetric: Metric = dummyMetric;

  // fillMetrics() {
  //     const figures = document.querySelector("div")?.querySelectorAll("figure")!

  //     if (!figures) return;
      
  //     for (const figure of figures) {
  //       let newMetric: Metric = {
  //         Label: "",
  //         Value: 0,
  //         Explanation: "",
  //         Operation: "",
  //         Metrics: [],
  //         Rating: -1,
  //       }
  //       const allPs = figure.querySelectorAll("p")!
  //       let count = 0;
  //       for (const p of allPs) {
  //         if (count % 4 === 0) {  // 0 mod 4 == 0
  //           newMetric.Label = p.innerHTML
  //         } else if (count % 4 === 1) {  // 1 mod 4 === 1
  //           newMetric.Value = Number(p.innerHTML)
  //         } else if (count % 4 === 2) {  // 2 mod 4 === 2
  //           newMetric.Explanation = p.innerHTML
  //         } else if (count % 4 === 3) {  // 3 mod 4 === 3  
  //           newMetric.Operation = p.innerHTML
  //         }
          
  //         count++;
  //       }
  //       this.metrics.push(newMetric)

  //     }

  //     this.requestUpdate()
  // }

  
  connectedCallback() {
    super.connectedCallback()
    // setTimeout(this.fillMetrics.bind(this), 0);
  }

  addNewMetric() {
    if (this.newMetric !== dummyMetric && !this.metrics.includes(this.newMetric)) {
      AddNewMetric(this.newMetric, getDocId())
      this.metrics.push(this.newMetric)
    }
  }

  constructor() {
    super()
  }

  render() {

    this.addNewMetric()

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
        
        ${this.metrics.map((metric: Metric) => {

          return html`
          <metric-component
          Label=${metric.label}
          Value=${metric.value}
          Explanation=${metric.explanation}
          >
          </metric-component>
          `

        })}
        

        


       </section>
    `
  }
}



