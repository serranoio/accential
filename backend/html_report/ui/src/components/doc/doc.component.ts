import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import { tabs } from '../../model/tabs';
import docCss from './doc.css';
import { AddMetricSteps, LabelValueSteps, dummyMetric } from '../../model/metric';

@customElement('doc-component')
export class DocComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [docCss] 

  @property()
  openView = false;

  @property()
  selectedTab = tabs[0];

  @property()
  // @ts-ignore
  frame: HTMLDivElement;

  @property()
  creatingMetricFromDocument = structuredClone(dummyMetric)

  @property()
  labelValueSteps = LabelValueSteps.Label;

  @property()
  addMetricSteps = AddMetricSteps.AddMetric

  constructor() {
    super()
  }

  getValue(): string {
    return this.creatingMetricFromDocument.value === -1 ?
    "" :
    String(this.creatingMetricFromDocument.value);
  }

  getLabel(): string {
    const label = this.creatingMetricFromDocument.label;

    return label.length > 21 ? label.slice(0,21) + "..." : String(label); 
  }

  render() {

    return html`
        ${this.frame}

        ${this.addMetricSteps !== AddMetricSteps.AddMetric ?

          html`<aside class="labels">

        <h3>Current Metric</h3>
        <figure class="current-metric">
        <p class="${this.labelValueSteps === LabelValueSteps.Label ? "on" : ""}"
        >Label <p>&nbsp;${this.getLabel()}</p></p>
        <p class="${this.labelValueSteps === LabelValueSteps.Value ? "on" : ""}"
        >Value <p>&nbsp;${this.getValue()}</p></p>
        <p class="${this.labelValueSteps === LabelValueSteps.Explanation ? "on" : ""}"   
        >Explanation </p>
        </figure>
        
        </aside>`
        : ""
      }
    `
  }
}



