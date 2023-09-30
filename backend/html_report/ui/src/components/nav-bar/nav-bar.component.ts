import { CSSResultGroup, LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

import navBarCss from './nav-bar.css';
import { tabs } from '../../model/tabs';
import { SetSelectedTab } from '../../model/events';
import { dummyDoc } from '../../model/statistics';


@customElement('nav-bar')
export class NavBarComponent extends LitElement {
  static styles?: CSSResultGroup | undefined = [navBarCss] 

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


    this.dispatchEvent(new CustomEvent(SetSelectedTab, {
      composed: true,
      bubbles: true,
      detail: {
        selectedTab: innerHTML.split(">")[1].trim(),
      }
    }))
  }

  render() {

    
    return html`
        <nav>
        <h3>${dummyDoc.name} Report</h3>

        <ul @click=${this.handleClick}>
            ${tabs.map((tab: string) => html`
            <li class="tab
            ${this.selectedTab === tab ? "selected-tab" : ""}
            ">
            <p>
            ${tab}
            </p>
            </li>`)}
        </ul>
        </nav>
    `
  }
}



