import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import { tabs } from '../../model/tabs';
import docCss from './doc.css';


@customElement('doc-component')
export class DocComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [docCss] 

  @property()
  openView = false;

  @property()
  selectedTab = tabs[0];

  constructor() {
    super()
  }

  handleClick(e: any) {
    if (!e.target.closest("li")) return;

    let target = e.target.nodeName === "LI" ? e.target.querySelector("p") : e.target;
    let innerHTML = target.innerHTML;

    this.selectedTab = innerHTML.split(">")[1].trim();
  }

  render() {
    return html`
       <section class="full">
       hello
       </section>
    `
  }
}



