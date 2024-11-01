'use client';

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const schema = z.object({
    companyName: z.string().min(1),
    hourlyWage: z.number().min(0),
    description: z.string().min(1)
})

type FormValues = z.infer<typeof schema>

export const Form = () => {
    const { register, formState:{errors}, handleSubmit} = useForm<FormValues>({
        resolver: zodResolver(schema)
    })

    const onSubmit = (data: FormValues) => {
        console.log(data);
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <input type="text" {...register("companyName")} />
            <span>{errors.companyName && <span>{ errors.companyName.message}</span>}</span>
            { }
            <button type="submit">送信</button>
        </form>
    )
}