import { RouteSectionProps } from '@solidjs/router';
import { Suspense } from 'solid-js';
import Sidebar from '~/components/containers/sidebar/sidebar';
import { Col, Grid } from '~/components/ui/grid';
import TopHeader from './top-header';
import Player from './player';

const Layout = (props: RouteSectionProps) => {
  return (
    <div class="relative">
      <Grid cols={1} colsMd={12} class="w-full">
        <Col spanMd={3} spanLg={2}>
          <Sidebar />
        </Col>
        <Col span={1} spanMd={9} spanLg={10}>
          <TopHeader />
          <Suspense>{props.children}</Suspense>
        </Col>
      </Grid>
      <Player />
    </div>
  );
};
export default Layout;
