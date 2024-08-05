# Dependency Injection

#design-pattern #dependency-injection

> Dependency Injection is simply when you have a piece of code which uses
> another piece of code and instead of using that piece of code directly it's
> _passed in_ instead.
> 
> When you pass something in to be used, we call it injection.

> Dependency Injection: Passing in an interface for what you are going to use.

## Scenario

We have a business app where users can chat with their coworkers. They can also
send pictures and files to each other. When a user sends something, the file
gets uploaded to our attachment service. The attachment service is responsible
for storing, retrieving and processing all attachments.

We're going to build up this whole service using dependency injection, and we'll
see what it enables us to do.

When a user sends a message with an attachment, the message text gets sent to
our standard chat service. We want people to receive their messages almost real
time. So this service is all about speed.

The attachment, on the other hand, gets uploaded to our attachment service.
There's an endpoint in out note service `/attachment/upload` that the app
connects to and uploads a file. The attachment gets stored on the disk
temporarily, process in a few ways we'll talk about and the uploaded to its
final destination.

The default storage location is an S3, a part of Amazon's Web Services. It's a
simple storage service that lets you put up files and pull them down.

We have some code here that takes the uploaded file and then uploads it to S3.

```typescript
class Storage {
    async upload(attachment: Attachment, aws_access_key_id: string, aws_secret_access_key: string): Promise<string> {
        const attachment_id = uuidv4()
        AWS.config.credentials = new AWS.Credentials(aws_asses_key_id, aws_secret_access_key)
        const s3 = new AWS.S3({apiVersion: '2006-03-01'})
        await this.upload_to_s3(s3, attachment_id, attachment.message_id, attachment.local_path, attachment.preview_local_path)
        
        return attachment_id
    }
    
    async upload_to_s3(
        s3: AWS.S3, attachment_id: string, message_id: string, 
        local_path: string, preview_local_path: string
    ): Promise<void> {
         const main_upload = s3.upload({
            Bucket: "uploads",
             Key: attachment_id,
             Metadata: {
                 message_id: message_id,
             },
             Body: fs.createReadStream(local_path),
         })
        
        const uploads = [main_upload.promise()]
        
        if (preview_local_path) {
            const preview_upload = s3.upload({
                Bucket: "uploads",
                Key: attachment_id,
                Metadata: {
                    message_id: message_id,
                },
                Body: fs.createReadStream(local_preview_path),
            })
            uploads.push(preview_upload.promise())
        }
    }
}
```

Unfortunately, simple and elegant doesn't always like to co-exist with business.
While S3 is nice and straightforward and most of our clients are okay with it,
we have a few firms that don't want us to permanently store their data. This
means we actually need our service to handle multiple storage locations. Then
depending on which company a user is from, we need to put their attachments in
the right place. Most of these picky clients just give us an SFTP server to
connect to, but one really wanted us to use WebDav.

Our first thought might be to simply extend out upload code with some if
statements, like below:


```typescript
enum Destination {
    AWS,
    SFTP,
    WebDav,
}

interface StorageParams {
    destination: Destination
    aws_access_key_id?: string
    aws_secret_access_key?: string
    sftp_host?: string
    sftp_port?: number
    sftp_username?: string
    sftp_private_key?: string
    webdav_url?: string
    webdav_authorization_key?: string
}

class Storage {
    async upload(attachment: Attachment, params: StorageParams): Promise<string> {
        const attachment_id = uuidv4()
        
        if (params.destination === Destination.AWS) {
            if (!params.aws_access_key_id || !params.aws_secret_access_key) {
                throw new Error("Missing aws configuration") 
            }

            AWS.config.credentials = new AWS.Credentials(aws_asses_key_id, aws_secret_access_key)
            const s3 = new AWS.S3({apiVersion: '2006-03-01'})
            await this.upload_to_s3(s3, attachment_id, attachment.message_id, attachment.local_path, attachment.preview_local_path)
        }
        
        if (params.destination === Destination.SFTP) {
            if (!params.sftp_host || !params.sftp_port || !params.stfp_username || !params.sftp_privat_key) {
                throw new Error("Missing sftp configuration")
            } 
            
            const options = {
                host: params.sftp_host,
                port: params.sftp_port,
                username: params.sftp_username,
                privateKey: params.sftp_private_key,
            }
            
            await this.upload_to_sftp(attachment, attachment_id, options)
        }
        
        if (params.destination === Destination.WebDav) {
            // similar checks for web dav 
            await this.upload_to_webdav(attachment, attachment_id, options)
        }

        return attachment_id
    }

    async upload_to_s3(
        s3: AWS.S3, attachment_id: string, message_id: string,
        local_path: string, preview_local_path: string
    ): Promise<void> {
        const main_upload = s3.upload({
            Bucket: "uploads",
            Key: attachment_id,
            Metadata: {
                message_id: message_id,
            },
            Body: fs.createReadStream(local_path),
        })

        const uploads = [main_upload.promise()]

        if (preview_local_path) {
            const preview_upload = s3.upload({
                Bucket: "uploads",
                Key: attachment_id,
                Metadata: {
                    message_id: message_id,
                },
                Body: fs.createReadStream(local_preview_path),
            })
            uploads.push(preview_upload.promise())
        }
    }
}

async function uploadCaller() {
    let params: StorageParams
    
    // this company want sftp
    if (this.storage_config.sftp_companies[this.company_id]) {
        const sftp = this.storage_config.sftp_companies[this.company_id] 
        params = {
            destination: Destination.SFTP,
            sftp_host: sftp.host,
            sftp_port: sftp.port,
            sftp_username: sftp.username,
            sftp_private_key: sftp.private_key,
        }
    } else if (this.storage_config.webdav_companies[this.company_id]) {
        const webdav = this.storage_config.webdav_companies[this.company_id]
        params = {
            destination: Destination.WebDav,
            webdav_url: webdav.url,
            webdav_authorization_key: webdav.authorization_key,
        }
    } else {
        params = {
            destination: Destination.AWS,
            aws_access_key_id: "",
            aws_secret_access_key: "",
        }
    }
    
    let storage = new Storage()
    return await storage.upload(attachment, params)
}
```

And then have the caller of the upload method pass in where to upload the file.
This is awkward for a few reasons.

First, this one class has a ton of responsibility, making the code pretty ugly.
The code for SFTP is intermingled with the code from AWS and WebDav, even though
they're pretty different.

There's a lot of paths the code can take, and that makes the code harder to
understand.

Second, using the class isn't very simple. We have this one upload function
which needs a bunch of info for where we're going to upload the attachment.

```typescript
interface StorageParams {
    destination: Destination
    // when destinaiton === AWS
    aws_access_key_id?: string
    // when destinaiton === AWS
    aws_secret_access_key?: string
    // when destinaiton === SFTP 
    sftp_host?: string
    // when destinaiton === SFTP 
    sftp_port?: number
    // when destinaiton === SFTP 
    sftp_username?: string
    // when destinaiton === SFTP 
    sftp_private_key?: string
    // when destinaiton === WebDav 
    webdav_url?: string
    // when destinaiton === WebDav 
    webdav_authorization_key?: string
}
```

But what info it needs is very different depending on where it's going. If it's
AWS, we only need the AWS config. If it's SFTP we only need the SFTP config.
If it's WebDav we only need its config. Furthermore, extending the
implementation will only make this more difficult/complex.

So we are forced to have a bunch of these optional variables that need to be
filled out in certain cases, and then comments to tell you which ones to fill
out. This makes it pretty easy for the caller to make a mistake.

Finally, the part of the code that actually calls upload `uploadCaller` needs
to have all of this destination specific context to perform the upload. But
really, at this phase, it just wants to upload. The part of the code that knows
best which company a user is from and can deduce where the file should go is
at the beginning of the request, where the `company_id` is specified.

### Let's use Dependency Injection instead

Let's create an interface that represents out attachment storage, which contains
a key upload method that does what the request handler want to do: Upload an
attachment.

```typescript
interface AttachmentStorage {
    /**
     * Store an attachment
     *
     * @param attachment the attachment to store
     * @return the attachment id
     */
    upload(attachment: Attachment): Promise<string>
    /**
     * Retrievet the attachment from the storage server
     * @param attachemtn_id the attachment id
     */
    download(attachment_id: string): Promise
}
```

Then we create three different implementations of the storage interface. The
configuration for each is passed into their constructor. There is no more
optional variables that sometimes need to be set. We require exactly these
values and get an error if you forget one.

```typescript
class SftpStorage implements AttachmentStorage {
    private readonly host: string
    private readonly port: number
    private readonly username: string
    private readonly private_key: Buffer
    private readonly client = new sftp()
    
    public constructor(host: string, port: number, username: string, private_key: Buffer) {
        this.host = host
        this.port = port
        this.username = username
        this.private_key = private_key
    }
    
    async upload(attachment: Attachment) { /*sftp upload implementation*/ return Promise.resolve("") }
    async download(attachment_id: string) { /*sftp upload implementation*/ return Promise.resolve("") }
}

// class S3Storage implements Storage { }

class WebDavStorage implements AttachmentStorage {
    private webbdav_client: WebDavClient
    public constructor(uri: string, authorization_key: string) {
        this.webdav_client = createClient(uri, {
            headers: {
                authorization: `Authorization ${authorization_key}` 
            } 
        })
    }
    
    async upload(attachment: Attachment) { /*webdav upload implementation*/ return Promise.resolve("") }
    async download(attachment_id: string) { /*webdav upload implementation*/ return Promise.resolve("") }
}
```

So now, once the user is authenticated, and we know which company they're from,
we create the storage that the request handler should use.

```typescript
class UploadReqeust {
    private readonly storage: AttachmentStorage
    
    constructor(storage: AttachmentStorage) {
        this.storage = storage 
    }
    
    async upload(attachment: Attachment, mime_type: string) {
        const scanner = new ThreatProtectScanner()
        
        if (await scanner.scan(attachment.local_path)) {
            throw Error("Virus scanner detected issue") 
        }
    }
}
```

```typescript
app.post("/attachment/upload", async (req, res) => {
    const user_id = req.user_id
    const company_id = req.company_id
    
    let storage: AttachmentStorage
    
    if (this.storage_config.sftp_companies[this.company_id]) {
        const config = this.storage_config.sftp_companies[this.company_id]
        storage = new SftpStorage(config.host, config.port, config.username, config.private_key)
    } else if (this.storage_config.webdav_companies[this.company_id]) {
        const config = this.storage_config.webdav_companies[this.company_id]
        storage = new WebDavStorage(config.uri, config.authorization_key)
    } else {
        storage = new AwsStorage(/*add params*/)
    }
    
    const file = req.file
    if (!file) {
        res.status(400).send({error: "no file"})
        return
    }
    
    try {
        const attachment: Attachment = {
            local_path: file.path,
            user_id: user_id,
            message_id: req.body.message_id,
        }
        const request = new UploadRequest(storage)
        const attachment_id = await request.upload(attachment)
        res.send({attachment_id: attachment_id})
    } catch(e) {
        res.status(500).send({error: e})
    }
}) 
```

That said, the construction code will be too complex to put here, so lets see
if we can clean this up. If we look at the input here, it's really just this
company configuration and the output is the storage which conforms to the
storage interface. So we can just move this out to a factory.

```typescript
class StorageFactory {
    static createStorage(company_id): AttachmentStorage {
        if (this.storage_config.sftp_companies[this.company_id]) {
            const config = this.storage_config.sftp_companies[this.company_id]
            storage = new SftpStorage(config.host, config.port, config.username, config.private_key)
        } else if (this.storage_config.webdav_companies[this.company_id]) {
            const config = this.storage_config.webdav_companies[this.company_id]
            storage = new WebDavStorage(config.uri, config.authorization_key)
        } else {
            storage = new AwsStorage(/*add params*/)
        }
        return storage
    }
}

app.post("/attachment/upload", async (req, res) => {
    const user_id = req.user_id
    const company_id = req.company_id

    let storage: AttachmentStorage = StorageFactory.createStorage(company_id)

    const file = req.file
    if (!file) {
        res.status(400).send({error: "no file"})
        return
    }

    try {
        const attachment: Attachment = {
            local_path: file.path,
            user_id: user_id,
            message_id: req.body.message_id,
        }
        const request = new UploadReqeust(storage)
        const attachment_id = await request.upload(attachment)
        res.send({attachment_id: attachment_id})
    } catch(e) {
        res.status(500).send({error: e})
    }
}) 
```

Great, we have used two patterns. One, the _factory pattern_, which creates or
instantiates an object/class/interface without us knowing the exact config
or implementation. Two, _dependency injection_ as the `AttachmentStorage`
interface allows us to switch the underlying implementation of how to handle
the upload and download.