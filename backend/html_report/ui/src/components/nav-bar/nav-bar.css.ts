import { css } from "lit";

export default css`

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

}

html {
    font-size: 62.5%;
}

nav {
    display: flex;
    background-color: var(--gray20);
    justify-content: space-between;
    align-items: center;
    padding: 0 2.4rem;
    font-size: 2.4rem;
    height: 70px; 
}

ul {
    align-self: end;
    display: flex;
    align-items: end;
    gap: .6rem;
    list-style: none;
    // height: 100%;
}

h3, .tab {
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    background-color: var(--gray22);
    padding: 1.2rem 2.4rem;
    align-self: end;
    color: var(--gray60);
}

.tab {
    display: block;
    border-bottom: 2px solid var(--gray25);
    height: 100%;
    cursor: pointer;
    transition: all .2s;
}

.selected-tab {
    color: var(--gray);
    border-bottom: 2px solid var(--secondary-3);
}


`;