import { CSSResultGroup, LitElement, TemplateResult, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import mainStyles from "./main.css"
import { Statistics, dummyStatistics } from './model/statistics';
import './components/nav-bar/nav-bar.component';
import "./components/doc/doc.component"
import "./components/distilled/distilled.component"
import { CreateMetric, Distilled, Doc } from './model/tabs';
import { SetSelectedTab } from './model/events';
import { AddMetricSteps, CreateMetricOptions, LabelValueSteps, Metric, dummyMetric } from './model/metric';

@customElement('main-component')
export class Main extends LitElement {
  static styles?: CSSResultGroup | undefined = [mainStyles] 

  @property()
  statistics: Statistics;

  @property()
  metrics: Metric[] = [structuredClone(dummyMetric)]

  @property()
  selectedTab = CreateMetric;

  @property()
  newMetric: Metric = dummyMetric;

  @property()
  chosenMethod = CreateMetricOptions.SetManually;

  setSelectedTab(e: any) {
    this.selectedTab = e.detail.selectedTab;
  }

  @property()
  // @ts-ignore
  frame: HTMLDivElement

  @property()
  inserted: boolean  = false;

  @property()
  addMetricSteps = AddMetricSteps.AddMetric;  

  @property()
  currentElementList: any = [];

  @property()
  creatingMetricFromDocument = structuredClone(dummyMetric)

  @property()
  labelValueSteps = LabelValueSteps.Label;

  @property()
  creatingMetricInputs: number = -1;

  giveStyles(element: HTMLElement) {
    element.style.boxShadow = "0 0 0 2px var(--info)";
    element.style.zIndex = "10";
  }
  
  removeStyles(element: HTMLElement) {
    element.style.boxShadow = "none";
    element.style.zIndex = "0";
  }

  extractContent(element: HTMLElement, advanceStep: boolean) {
    if (this.labelValueSteps === LabelValueSteps.Label) {
      this.creatingMetricFromDocument.Label = element.innerText;
      if (advanceStep) {
        this.labelValueSteps = LabelValueSteps.Value
        this.metrics[this.creatingMetricInputs].Label = this.creatingMetricFromDocument.Label
        this.metrics = structuredClone(this.metrics)
      }
    } else if (this.labelValueSteps === LabelValueSteps.Value) {
      const num = element.innerText.replace(/\,/g,'');
      
      this.creatingMetricFromDocument.Value = Number(num);
      if (advanceStep) {
        this.labelValueSteps = LabelValueSteps.Explanation
        this.selectedTab = CreateMetric
        this.metrics[this.creatingMetricInputs].Value = this.creatingMetricFromDocument.Value
        this.metrics = structuredClone(this.metrics)
      }
    }
    this.creatingMetricFromDocument = structuredClone(this.creatingMetricFromDocument)
  }


  onDocHover() {
    const doc = document.querySelector(".doc")!
  
    doc?.remove()

    this.frame = doc as HTMLDivElement;

    // this.frame.onclick = (e) => {
    //   const element = e.target as HTMLElement;
        
    //   if (element.nodeName !== "P" &&
    //     element.nodeName !== "a") {
    //       return
    //   }

    //   this.giveStyles(element)
    //   this.extractContent(element, true)

    //   this.currentElementList.push({
    //     reference: element,
    //     isHovered: false,
    //     isClicked: true,
    //     pos: this.creatingMetricInputs,
    //   })

    // }

    this.frame.onmouseover = (e) => {
      const element = e.target as HTMLElement;
        
      if (element.nodeName !== "P" &&
        element.nodeName !== "a") {
          return
      }

      this.giveStyles(element)
      this.extractContent(element, false)

      if (!this.currentElementList.map((el: any) => el.reference).includes(element)) {
        this.currentElementList.push({
          reference: element,
          isHovered: true,
          isClicked: false,
          pos: this.creatingMetricInputs,
        })
      }
      

      this.currentElementList.forEach((curElement: any, _: number) => {
          curElement.reference.onmouseleave = () => {
            if (curElement.isClicked) return;
            this.removeStyles(curElement.reference)
            this.currentElementList = this.currentElementList.filter((elementToRemove: any) => 
            elementToRemove.reference !== curElement.reference)
          }

          curElement.reference.onclick = () => {
            curElement.isClicked = true;
            curElement.pos = this.creatingMetricInputs;
            this.giveStyles(element)
            this.extractContent(element, true)
          }
      })

      
    }


    this.inserted = true;
  }
  connectedCallback(): void {
    super.connectedCallback()

    if (this.inserted) return; 

    setTimeout(this.onDocHover.bind(this), 0);
  }

  removeSelectedElements() {

    this.currentElementList = this.currentElementList.filter(
      (elementsToKeep: any, _: number) => {
       if (elementsToKeep.pos !== this.creatingMetricInputs) {
         return true;
        } 

      this.removeStyles(elementsToKeep.reference)
       return false;
    })
  }


  changeMethod(e: any) {
    this.chosenMethod = e.detail.method.split(">")[1];

    if (this.chosenMethod === CreateMetricOptions.FromDocument) {
      this.selectedTab = Doc
      this.creatingMetricFromDocument = structuredClone(dummyMetric)
      this.labelValueSteps = LabelValueSteps.Label
      this.removeSelectedElements()
    }
  }

  changeStep(e: any) {
    this.addMetricSteps = e.detail.method;
  }

  setCreatingMetricInputs(e: any) {
    this.creatingMetricInputs = e.detail.creatingMetricInputs
  }

  updateMetrics(e: any) {
    this.metrics = e.detail.metrics;
  }

  addNewMetric(e: any) {
    this.newMetric = e.detail.metric;

    this.requestUpdate()
  }

  constructor() {
    super()
    this.statistics = dummyStatistics;

    document.addEventListener(SetSelectedTab, this.setSelectedTab.bind(this))
    document.addEventListener(AddMetricSteps.AddingMetric, this.changeMethod.bind(this))
    document.addEventListener("MetricSteps", this.changeStep.bind(this))
    document.addEventListener("creating-metric-inputs", this.setCreatingMetricInputs.bind(this))
    document.addEventListener("add-new-metric", this.addNewMetric.bind(this))
    document.addEventListener("update-metrics", this.updateMetrics.bind(this))
  }

  getMain() {
    let tab: TemplateResult = html``;

    if (this.selectedTab === Doc) {

      // tab = html` <slot name="doc"></slot>`
      
      tab =  html`<doc-component
      slot="doc"
      .frame="${this.frame}"
      .creatingMetricFromDocument=${this.creatingMetricFromDocument}
      .labelValueSteps=${this.labelValueSteps}
      >
      </doc-component>`
    } else if (this.selectedTab === Distilled) {
      // tab = html` <slot name="distilled"></slot>`
      tab = html`
      <distilled-component
      .statistics=${this.statistics}
      .newMetric=${this.newMetric}
      slot="distilled"
      ></distilled-component>`
    } else if (this.selectedTab === CreateMetric) {
      // tab = html` <slot name="create-metric"></slot>`
      tab = html`
      <create-metric-component
      slot="create"
      .addMetricSteps=${this.addMetricSteps}
      .chosenMethod=${this.chosenMethod}
      .creatingMetricInputs=${this.creatingMetricInputs}
      .metrics=${this.metrics}
      >
      </create-metric-component>
      `
    }

    return tab;
  }

  getStatus(): string {
    if (this.chosenMethod === CreateMetricOptions.FromDocument) {
      return "Document"
    } else if (this.chosenMethod === CreateMetricOptions.SetManually) {
      return "Manual"
    } else {
      return "Outside Source"
    }
  }


  render() {
    return html`
    <slot name="distilled"></slot>
    <main>
    <nav-bar
    .selectedTab=${this.selectedTab}
    >
    </nav-bar>
    
    ${this.getMain()}

    ${this.addMetricSteps === AddMetricSteps.AddingMetric ? html`
    <button class="status">Status: ${this.getStatus()}</button>
    ` : ""
  }
    </main>
    `
  }
}



