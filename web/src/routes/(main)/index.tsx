import { RouteSectionProps } from '@solidjs/router';
import Sidebar from '~/components/containers/sidebar/sidebar';
import { Col, Grid } from '~/components/ui/grid';

export default function MainLayout(props: RouteSectionProps) {
  return (
    <main>
      <Grid cols={1} colsMd={12} class="w-full">
        <Col spanMd={3} spanLg={2}>
          <Sidebar />
        </Col>
        <Col span={1} spanMd={9} spanLg={10}>
          {props.children}
        </Col>
      </Grid>
    </main>
  );
}
