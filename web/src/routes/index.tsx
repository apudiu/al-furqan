import { RouteSectionProps } from '@solidjs/router';

const Homepage = (props: RouteSectionProps) => {
  return <div>Homepage {props.location.pathname}</div>;
};
export default Homepage;
