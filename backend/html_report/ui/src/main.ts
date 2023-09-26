import { CSSResultGroup, LitElement, TemplateResult, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import mainStyles from "./main.css"
import { Statistics, dummyStatistics } from './model/statistics';
import './components/nav-bar/nav-bar.component';
import "./components/doc/doc.component"
import "./components/distilled/distilled.component"
import { CreateMetric, Distilled, Doc } from './model/tabs';
import { SetSelectedTab } from './model/events';

@customElement('main-component')
export class Main extends LitElement {
  static styles?: CSSResultGroup | undefined = [mainStyles] 

  @property()
  statistics: Statistics;

  @property()
  selectedTab = Distilled;

  setSelectedTab(e: any) {
    this.selectedTab = e.detail.selectedTab;
  }

  constructor() {
    super()
    this.statistics = dummyStatistics;


    document.addEventListener(SetSelectedTab, this.setSelectedTab.bind(this))
  }


  getMain() {
    let tab: TemplateResult = html``;

    if (this.selectedTab === Doc) {
      tab =  html`<doc-component>
      </doc-component>`
    } else if (this.selectedTab === Distilled) {
      tab = html`
      <distilled-component
      .statistics=${this.statistics}
      ></distilled-component>`
    } else if (this.selectedTab === CreateMetric) {
      tab = html`
      <create-metric-component>
      </create-metric-component>
      `
    }

    return tab;
  }

  render() {



    return html`
    <main>
    <nav-bar
    .selectedTab=${this.selectedTab}
    >
    </nav-bar>

    ${this.getMain()}
  


    
    </main>
    `
  }
}



