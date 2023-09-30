import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import { tabs } from '../../model/tabs';
import docCss from './doc.css';
import { LabelValueSteps, dummyMetric } from '../../model/metric';

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

  constructor() {
    super()
  }

  getValue(): string {
    return this.creatingMetricFromDocument.Value === -1 ?
    "" :
    String(this.creatingMetricFromDocument.Value);
  }

  getLabel(): string {
    const label = this.creatingMetricFromDocument.Label;

    return label.length > 8 ? label.slice(0,8) + "..." : String(label); 
  }

  render() {

    return html`
        ${this.frame}

        <aside class="labels">

        <h3>Current Metric</h3>
          <figure class="current-metric">
          <p class="${this.labelValueSteps === LabelValueSteps.Label ? "on" : ""}"
          >Label: ${this.getLabel()}</p>
          <p class="${this.labelValueSteps === LabelValueSteps.Value ? "on" : ""}"
          >Value: ${this.getValue()}</p>
          <p class="${this.labelValueSteps === LabelValueSteps.Explanation ? "on" : ""}"   
          >Explanation: </p>
          </figure>
          
        </aside>
    `
  }
}



