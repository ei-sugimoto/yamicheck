import { Stack, Typography } from "@mui/material";

export default function Header() {
    return (
        <Stack direction={"row"} gap={3} >
            <Typography variant="h5" gutterBottom>
                 闇バイトチェッカー
            </Typography>
            <Typography variant="inherit" gutterBottom>
                このサイトは、闇バイトをチェックするためのサイトです。
            </Typography>
       </Stack>
    )
}