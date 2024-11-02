import { zodResolver } from "@hookform/resolvers/zod";
import { Box, Button, Stack, TextField } from "@mui/material";
import { Controller, useForm } from "react-hook-form";
import { z } from "zod";

const schema = z.object({
    companyName: z.string().min(1, { message: "会社名を入力してください" }),
    hourlyWage: z.number(),
    description: z.string().min(1, { message: "求人内容を入力してください" }),
})

type FormValues = z.infer<typeof schema>;

export default function Form() {
    const { control, handleSubmit, formState: { errors } } = useForm<FormValues>({
        resolver: zodResolver(schema),
    });

    const onSubmit = (data: FormValues) => {
        console.log(data);
    }

    return (
        <Stack gap={3}>
            <Box component="form" onSubmit={handleSubmit(onSubmit)}>
                <Stack gap={3}>
                    <Controller
                        name="companyName"
                        control={control}
                        defaultValue=""
                        render={({ field }) => (
                            <TextField
                                {...field}
                                label="会社名"
                                error={!!errors.companyName}
                                helperText={errors.companyName?.message}
                            />
                        )}
                    />
                    <Controller
                        name="hourlyWage"
                        control={control}
                        defaultValue={0}
                        render={({ field }) => (
                            <TextField
                                {...field}
                                label="時給"
                                type="number"
                                error={!!errors.hourlyWage}
                                helperText={errors.hourlyWage?.message}
                            />
                        )}
                    />
                    <Controller
                        name="description"
                        control={control}
                        defaultValue=""
                        render={({ field }) => (
                            <TextField
                                {...field}
                                multiline
                                type=""
                                label="求人内容"
                                error={!!errors.description}
                                helperText={errors.description?.message}
                            />
                        )}
                    />
                    <Button type="submit" variant="contained">Submit</Button>
                </Stack>
            </Box>
        </Stack>
    );
}