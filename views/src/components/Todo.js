import Checkbox from '@mui/material/Checkbox';
import { Link, Grid } from '@mui/material';

function Todo() {
  return (
    <div>
      <Grid container spacing={1}>
        <Grid item xs={3} justifyContent={"center"} alignContent={"center"}>
          <Checkbox />
          <Link >
            <span>タイトルを書く</span>
          </Link>
        </Grid>
        <Grid item xs={2}>
          <div>20XX年○月×日</div>
        </Grid>
      </Grid>
    </div>
  )
}

export default Todo;
