import { css } from "lit";

export default css`
html {
  font-size: 62.5%;
}
* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

main {
  margin: 0 auto;
  max-width: 140rem;

  position: relative;
}

h1 {
  text-align: center;
  font-size: var(--fontSize1);
}


.status {
  position: absolute;
  right: 0%;
  transform: translateX(50%);
  bottom: 5%;
  background-color: var(--gray30);
  padding: var(--spacingHalf) var(--spacing);
  font-size: 2rem;
  border-radius: 10px;
  color: var(--80);
  border: none;
}

`