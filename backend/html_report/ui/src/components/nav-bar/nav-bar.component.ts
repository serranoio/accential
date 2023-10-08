import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import navBarCss from './nav-bar.css';
import { tabs } from '../../model/tabs';
import { SetSelectedTab } from '../../model/events';
import { AddMetricSteps } from '../../model/metric';
import { ChangeDocumentName, FetchName } from '../../model/api';
import { getDocId } from '../../model/util';


@customElement('nav-bar')
export class NavBarComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [navBarCss] 

  @property()
  openView = false;

  @property()
  selectedTab = tabs[0];

  @property()
  addMetricSteps = AddMetricSteps.AddMetric;

  @property()
  finishEditing = false;

  @property()
  clicked = false;

  @property()
  nameDocument = "";

  constructor() {
    super()
  }

  async fetchName() {

    try {
      this.nameDocument = await FetchName(getDocId());

    } catch(err) {

    }

  }

  connectedCallback(): void {
    super.connectedCallback();

    this.fetchName();
  }

  handleClick(e: any) {
    if (this.addMetricSteps !== AddMetricSteps.AddMetric) {
      this.clicked = true;
      return;
    }

    if (!e.target.closest("li")) return;

    let target = e.target.nodeName === "LI" ? e.target.querySelector("p") : e.target;
    let innerHTML = target.innerHTML;

    
    this.selectedTab = innerHTML.split(">")[1].trim();


    this.dispatchEvent(new CustomEvent(SetSelectedTab, {
      composed: true,
      bubbles: true,
      detail: {
        selectedTab: innerHTML.split(">")[1].trim(),
      }
    }))
  }

  reportName(e: any) {
    this.nameDocument = e.target.value; 
  }

  changedName(e: any) {
    this.nameDocument = e.target.value;

    // @ts-ignore
    ChangeDocumentName(getDocId(),
     this.nameDocument)

  }

  render() {
    if (this.addMetricSteps !== AddMetricSteps.AddMetric) {
      this.finishEditing = true;
    } else {
      this.clicked = false;
      this.finishEditing = false;
    }
    
    return html`
        <nav>
        <h3><input type="text"
          value=${this.nameDocument}
         @input=${this.reportName}
         @change=${this.changedName}
         placeholder="Fill Name!"
         class=${this.nameDocument === "" ? "fill-in-name" : ""}
         /></h3>

        <ul @click=${this.handleClick}>
            ${tabs.map((tab: string) => html`
            <li class="tab
            ${this.selectedTab === tab ? "selected-tab" : ""}
            ">
            <p>
            ${tab}
            
            </p>
            
            </li>`)}
            <span class="
            popup
            ${this.clicked ? "brighten" : ""}
            "
            >
            finish editing</span>
        </ul>
        </nav>
    `
  }
}



