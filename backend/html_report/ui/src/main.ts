import { CSSResultGroup, LitElement, TemplateResult, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import mainStyles from "./main.css"
import './components/nav-bar/nav-bar.component';
import "./components/doc/doc.component"
import "./components/distilled/distilled.component"
import { CreateMetric, Distilled, Doc } from './model/tabs';
import { AddNewMetricEvent, ChangeMetricSteps, DeleteMetric, EditMetric, SetCurrentInput, SetSelectedTab, UpdateSubmetrics, UseMetric } from './model/events';
import { AddMetricSteps, CreateMetricOptions, LabelValueSteps, Metric, Submetric, dummyMetric, dummySubmetric } from './model/metric';
import { AddNewMetric, DeleteMetricDb, GetAllMetrics } from './model/api';
import { getDocId } from './model/util';

@customElement('main-component')
export class Main extends LitElement {
  static styles?: CSSResultGroup | undefined = [mainStyles] 

  @property()
  creatingMetrics: Submetric[] = [structuredClone(dummySubmetric)]
  @property()
  creatingMetricsMain: Metric = structuredClone(dummyMetric)

  @property()
  selectedTab = Doc;  
  
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
  // any time we are creating a new metric from document, we populate this guy
  @property()
  creatingMetricFromDocument = structuredClone(dummyMetric)

  @property()
  labelValueSteps = LabelValueSteps.Label;
  // counter to describe which set of inputs we are selecting
  @property()
  creatingMetricInputs: number = -1;

  @property()
  position: number = -1;

  giveStyles(element: HTMLElement) {
    element.style.boxShadow = "inset 0 0 0 2px var(--info)";
    element.style.zIndex = "10";
  }
  
  removeStyles(element: HTMLElement) {
    element.style.boxShadow = "none";
    element.style.zIndex = "0";
  }

  extractContent(element: HTMLElement, advanceStep: boolean) {
    if (this.labelValueSteps === LabelValueSteps.Label) {
      this.creatingMetricFromDocument.label = element.innerText;
      if (advanceStep) {
        this.labelValueSteps = LabelValueSteps.Value
        this.creatingMetrics[this.creatingMetricInputs].label = this.creatingMetricFromDocument.label
        this.creatingMetrics = structuredClone(this.creatingMetrics)
      }
    } else if (this.labelValueSteps === LabelValueSteps.Value) {
      let num = element.innerText.replace(/\,/g,'');  // tbh, idk what this is
      
      if (num[0] === "(" || num[num.length-1] === ")") {
        num = num.replace("(", "-")
        num = num.replace(")", "")
      }

      this.creatingMetricFromDocument.value = Number(num);
      if (advanceStep) {
        this.labelValueSteps = LabelValueSteps.Explanation
        this.selectedTab = CreateMetric
        this.creatingMetrics[this.creatingMetricInputs].value = this.creatingMetricFromDocument.value
        this.creatingMetrics = structuredClone(this.creatingMetrics)
      }
    }
    this.creatingMetricFromDocument = structuredClone(this.creatingMetricFromDocument)
  }

  onDocHover() {
    const doc = document.querySelector(".doc")!
  
    doc?.remove()

    this.frame = doc as HTMLDivElement;

    this.frame.onmouseover = (e) => {
      let element = e.target as HTMLElement;

      element = element.closest("td")!;
      if (!element || element.nodeName !== "TD" || this.addMetricSteps === AddMetricSteps.AddMetric) {
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

  @property()
  // @ts-ignore
  metrics: Metric[] = structuredClone(dummyMetric)

  async getAllMetrics() {
    try {

      this.metrics = await GetAllMetrics(getDocId())
    } catch(err) {
      this.metrics = [dummyMetric]
    }
  }

  connectedCallback(): void {
    super.connectedCallback()

    if (this.inserted) return; 

    setTimeout(this.onDocHover.bind(this), 0);

    this.getAllMetrics()
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
    } else if (this.chosenMethod === CreateMetricOptions.FromOthers) {
      if (this.metrics.length > 0) {
        this.selectedTab = Distilled
      }
    }
  }

  changeStep(e: any) {
    this.addMetricSteps = e.detail.method;

    if (this.addMetricSteps === AddMetricSteps.AddingMetric) {
      this.chosenMethod = CreateMetricOptions.SetManually
    } else if (this.addMetricSteps === AddMetricSteps.AddMetric) {
      this.creatingMetricFromDocument = structuredClone(dummyMetric);
      this.labelValueSteps = LabelValueSteps.Label
    }
  }

  setCreatingMetricInputs(e: any) {
    this.creatingMetricInputs = e.detail.creatingMetricInputs
  }

  updateSubmetrics(e: any) {
    this.creatingMetrics = e.detail.metrics;
  }

  addNewMetric(e: any) {
    let newMetric = e.detail.metric
    AddNewMetric(newMetric, getDocId())

    // once we add new metric, we have to reset everything
    this.creatingMetricsMain = structuredClone(dummyMetric) 
    this.creatingMetrics = [structuredClone(dummySubmetric)]
    
    // if we are editing an old one, we have to update the old one
    if (this.position !== -1) {
      this.metrics[this.position] = newMetric;  
      this.position = -1;
    } else {  // if we are ADDING, then add
      this.metrics.push(newMetric)
    }
 
    this.addMetricSteps = AddMetricSteps.AddMetric;
    this.chosenMethod = CreateMetricOptions.SetManually;

    this.requestUpdate()
  }

  getUseMetric(e: any) {
    let metric = e.detail.metric;
    metric.operation = ""  // becaue it would show up as (a/a)
    this.creatingMetrics[this.creatingMetricInputs] = metric
    this.selectedTab = CreateMetric
  }

  getEditMetric(e: any) {
    this.creatingMetrics = e.detail.metric.submetric;
    this.creatingMetricsMain = e.detail.metric;
    this.selectedTab = CreateMetric;
    this.position = e.detail.position;
  }

  getDeleteMetric(e: any) {
    let deletingMetric = e.detail.metric;
    this.metrics = this.metrics.filter((metric: Metric) => metric !== deletingMetric);
    DeleteMetricDb(deletingMetric, getDocId());
  }

  constructor() {
    super()

    document.addEventListener(SetSelectedTab, this.setSelectedTab.bind(this))
    document.addEventListener(AddMetricSteps.AddingMetric, this.changeMethod.bind(this))
    document.addEventListener(ChangeMetricSteps, this.changeStep.bind(this))
    document.addEventListener(SetCurrentInput, this.setCreatingMetricInputs.bind(this))
    document.addEventListener(AddNewMetricEvent, this.addNewMetric.bind(this))
    document.addEventListener(UpdateSubmetrics, this.updateSubmetrics.bind(this))
    document.addEventListener(UseMetric, this.getUseMetric.bind(this))
    document.addEventListener(EditMetric, this.getEditMetric.bind(this))
    document.addEventListener(DeleteMetric, this.getDeleteMetric.bind(this))
  }

  getMain() {
    let tab: TemplateResult = html``;

    if (this.selectedTab === Doc) {
      tab =  html`<doc-component
      slot="doc"
      .frame="${this.frame}"
      .creatingMetricFromDocument=${this.creatingMetricFromDocument}
      .labelValueSteps=${this.labelValueSteps}
      .addMetricSteps=${this.addMetricSteps}
      >
      </doc-component>`
    } else if (this.selectedTab === Distilled) {
      tab = html`
      <distilled-component
      .newMetric=${this.newMetric}
      .chosenMethod=${this.chosenMethod}
      slot="distilled"
      .metrics=${this.metrics}
      ></distilled-component>`
    } else if (this.selectedTab === CreateMetric) {
      tab = html`
      <create-metric-component
      slot="create"
      .addMetricSteps=${this.addMetricSteps}
      .chosenMethod=${this.chosenMethod}
      .creatingMetricInputs=${this.creatingMetricInputs}
      // main maps to metric, metric maps to submetric
      .metric=${this.creatingMetricsMain}
      .submetrics=${this.creatingMetrics}
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
    } else if (this.chosenMethod === CreateMetricOptions.FromOthers) {
      return "From Others"
    } else {
      return "From outside source"
    }
  }

  render() {
    return html`
    <slot name="distilled"></slot>
    <main>
      <nav-bar
      .selectedTab=${this.selectedTab}
      .addMetricSteps=${this.addMetricSteps}
      ></nav-bar>
      ${this.getMain()}
      ${this.addMetricSteps === AddMetricSteps.AddingMetric ? html`
      <button class="status">Status: ${this.getStatus()}</button>
      ` : ""
      }
    </main>
    `
  }
}



