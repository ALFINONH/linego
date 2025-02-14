U
    �^�^>  �                   @   s�  d dl mZmZmZmZmZ d dlmZ d dlm	Z	 d dl
Z
d dlZddlT d dl mZ d dlmZ g ZG d	d
� d
e�ZG dd� de�ZG dd� dee�ZG dd� de�Ze�e� ddejdedgdffe_G dd� de�Ze�e� d ejdedgdfdejdedgdffe_G dd� de�Ze�e� ddejdedgdffe_G dd� de�Ze�e� d ejde dgdfdejdedgdffe_e	e� [dS )�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolException)�fix_specN�   )�*)�
TProcessor)�
TTransportc                   @   s   e Zd Zdd� Zdd� ZdS )�Ifacec                 C   s   dS ��1
        Parameters:
         - request

        N� ��self�requestr   r   �</storage/sdcard0/bot1/flex/curve/LoginPermitNoticeService.py�checkQrCodeVerified   s    zIface.checkQrCodeVerifiedc                 C   s   dS r   r   r   r   r   r   �checkPinCodeVerified   s    zIface.checkPinCodeVerifiedN)�__name__�
__module__�__qualname__r   r   r   r   r   r   r      s   r   c                   @   sF   e Zd Zddd�Zdd� Zdd� Zdd	� Zd
d� Zdd� Zdd� Z	dS )�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   �iprot�oprotr   r   r   �__init__(   s    zClient.__init__c                 C   s   | � |� | �� S �r   )�send_checkQrCodeVerified�recv_checkQrCodeVerifiedr   r   r   r   r   .   s    
zClient.checkQrCodeVerifiedc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   �writeMessageBeginr   �CALLr   �checkQrCodeVerified_argsr   �write�writeMessageEnd�trans�flush�r   r   �argsr   r   r   r"   7   s    
zClient.send_checkQrCodeVerifiedc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz*checkQrCodeVerified failed: unknown result)r   �readMessageBeginr   �	EXCEPTIONr   �read�readMessageEnd�checkQrCodeVerified_result�success�e�MISSING_RESULT�r   r   �fname�mtype�rseqid�x�resultr   r   r   r#   ?   s    




zClient.recv_checkQrCodeVerifiedc                 C   s   | � |� | �� S r!   )�send_checkPinCodeVerified�recv_checkPinCodeVerifiedr   r   r   r   r   P   s    
zClient.checkPinCodeVerifiedc                 C   sF   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  d S )Nr   )r   r$   r   r%   r   �checkPinCodeVerified_argsr   r'   r(   r)   r*   r+   r   r   r   r;   Y   s    
z Client.send_checkPinCodeVerifiedc                 C   s�   | j }|�� \}}}|tjkr:t� }|�|� |��  |�t� }|�|� |��  |jd k	rb|jS |j	d k	rr|j	�ttj
d��d S )Nz+checkPinCodeVerified failed: unknown result)r   r-   r   r.   r   r/   r0   �checkPinCodeVerified_resultr2   r3   r4   r5   r   r   r   r<   a   s    




z Client.recv_checkPinCodeVerified)N)
r   r   r   r    r   r"   r#   r   r;   r<   r   r   r   r   r   '   s   
		r   c                   @   s4   e Zd Zdd� Zdd� Zdd� Zdd� Zd	d
� ZdS )�	Processorc                 C   s.   || _ i | _tj| jd< tj| jd< d | _d S )Nr   r   )�_handler�_processMapr?   �process_checkQrCodeVerified�process_checkPinCodeVerified�_on_message_begin)r   �handlerr   r   r   r    t   s
    zProcessor.__init__c                 C   s
   || _ d S �N)rD   )r   �funcr   r   r   �on_message_begin{   s    zProcessor.on_message_beginc                 C   s�   |� � \}}}| jr"| �|||� || jkr�|�tj� |��  ttjd| �}|�	|t
j|� |�|� |��  |j��  d S | j| | |||� dS )NzUnknown function %sT)r-   rD   rA   �skipr   �STRUCTr0   r   �UNKNOWN_METHODr$   r   r.   r'   r(   r)   r*   )r   r   r   �name�type�seqidr9   r   r   r   �process~   s    


zProcessor.processc           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )N�!TApplication exception in handler�Unexpected exception in handler�Internal errorr   )r&   r/   r0   r1   r@   r   r   r2   r   �REPLYr   �TTransportException�SecondaryQrCodeExceptionr3   r   �logging�	exceptionr.   �	Exception�INTERNAL_ERRORr$   r'   r(   r)   r*   �	r   rN   r   r   r,   r:   �msg_typer3   �exr   r   r   rB   �   s0    




z%Processor.process_checkQrCodeVerifiedc           	   
   C   s  t � }|�|� |��  t� }z| j�|j�|_tj	}W n� t
jk
rP   � Y n� tk
r| } ztj	}||_W 5 d }~X Y nd tk
r� } zt�d� tj}|}W 5 d }~X Y n0 tk
r�   t�d� tj}ttjd�}Y nX |�d||� |�|� |��  |j��  d S )NrP   rQ   rR   r   )r=   r/   r0   r>   r@   r   r   r2   r   rS   r   rT   rU   r3   r   rV   rW   r.   rX   rY   r$   r'   r(   r)   r*   rZ   r   r   r   rC   �   s0    




z&Processor.process_checkPinCodeVerifiedN)r   r   r   r    rH   rO   rB   rC   r   r   r   r   r?   s   s
   r?   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r&   �%
    Attributes:
     - request

    Nc                 C   s
   || _ d S rF   �r   r   r   r   r   r    �   s    z!checkQrCodeVerified_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr	   )�_fast_decode�
isinstancer)   r   �CReadableTransport�thrift_spec�	__class__�readStructBegin�readFieldBeginr   �STOPrJ   �&LoginQrCode_CheckQrCodeVerifiedRequestr   r/   rI   �readFieldEnd�readStructEnd�r   r   r6   �ftype�fidr   r   r   r/   �   s    "



zcheckQrCodeVerified_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr&   r   r	   ��_fast_encoderc   r)   r'   rd   �writeStructBeginr   �writeFieldBeginr   rJ   �writeFieldEnd�writeFieldStop�writeStructEnd�r   r   r   r   r   r'   �   s    

zcheckQrCodeVerified_args.writec                 C   s   d S rF   r   �r   r   r   r   �validate�   s    z!checkQrCodeVerified_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS �z%s=%rr   ��.0�key�valuer   r   r   �
<listcomp>�   s   �z5checkQrCodeVerified_args.__repr__.<locals>.<listcomp>�%s(%s)�, ��__dict__�itemsrd   r   �join�r   �Lr   r   r   �__repr__�   s    �z!checkQrCodeVerified_args.__repr__c                 C   s   t || j�o| j|jkS rF   �ra   rd   r�   �r   �otherr   r   r   �__eq__�   s    zcheckQrCodeVerified_args.__eq__c                 C   s
   | |k S rF   r   r�   r   r   r   �__ne__�   s    zcheckQrCodeVerified_args.__ne__)N�r   r   r   �__doc__r    r/   r'   rw   r�   r�   r�   r   r   r   r   r&   �   s   
r&   r   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r1   �.
    Attributes:
     - success
     - e

    Nc                 C   s   || _ || _d S rF   �r2   r3   �r   r2   r3   r   r   r   r      s    z#checkQrCodeVerified_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S �Nr   r	   )r`   ra   r)   r   rb   rc   rd   re   rf   r   rg   rJ   �'LoginQrCode_CheckQrCodeVerifiedResponser2   r/   rI   rU   r3   ri   rj   rk   r   r   r   r/     s(    "




zcheckQrCodeVerified_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr1   r2   r   r3   r	   �ro   rc   r)   r'   rd   rp   r2   rq   r   rJ   rr   r3   rs   rt   ru   r   r   r   r'   ,  s    


z checkQrCodeVerified_result.writec                 C   s   d S rF   r   rv   r   r   r   rw   <  s    z#checkQrCodeVerified_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS rx   r   ry   r   r   r   r}   @  s   �z7checkQrCodeVerified_result.__repr__.<locals>.<listcomp>r~   r   r�   r�   r   r   r   r�   ?  s    �z#checkQrCodeVerified_result.__repr__c                 C   s   t || j�o| j|jkS rF   r�   r�   r   r   r   r�   D  s    z!checkQrCodeVerified_result.__eq__c                 C   s
   | |k S rF   r   r�   r   r   r   r�   G  s    z!checkQrCodeVerified_result.__ne__)NNr�   r   r   r   r   r1     s   
r1   r2   r3   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r=   r]   Nc                 C   s
   || _ d S rF   r^   r   r   r   r   r    X  s    z"checkPinCodeVerified_args.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r_   )r`   ra   r)   r   rb   rc   rd   re   rf   r   rg   rJ   �'LoginQrCode_CheckPinCodeVerifiedRequestr   r/   rI   ri   rj   rk   r   r   r   r/   [  s    "



zcheckPinCodeVerified_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr=   r   r	   rn   ru   r   r   r   r'   o  s    

zcheckPinCodeVerified_args.writec                 C   s   d S rF   r   rv   r   r   r   rw   {  s    z"checkPinCodeVerified_args.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS rx   r   ry   r   r   r   r}     s   �z6checkPinCodeVerified_args.__repr__.<locals>.<listcomp>r~   r   r�   r�   r   r   r   r�   ~  s    �z"checkPinCodeVerified_args.__repr__c                 C   s   t || j�o| j|jkS rF   r�   r�   r   r   r   r�   �  s    z checkPinCodeVerified_args.__eq__c                 C   s
   | |k S rF   r   r�   r   r   r   r�   �  s    z checkPinCodeVerified_args.__ne__)Nr�   r   r   r   r   r=   P  s   
r=   c                   @   sJ   e Zd ZdZddd�Zdd� Zdd� Zd	d
� Zdd� Zdd� Z	dd� Z
dS )r>   r�   Nc                 C   s   || _ || _d S rF   r�   r�   r   r   r   r    �  s    z$checkPinCodeVerified_result.__init__c                 C   s�   |j d k	r<t|jtj�r<| jd k	r<|� | || j| jg� d S |��  |�� \}}}|t	j
kr^q�|dkr�|t	jkr�t� | _| j�|� q�|�|� n>|dkr�|t	jkr�t� | _| j�|� q�|�|� n
|�|� |��  qD|��  d S r�   )r`   ra   r)   r   rb   rc   rd   re   rf   r   rg   rJ   �(LoginQrCode_CheckPinCodeVerifiedResponser2   r/   rI   rU   r3   ri   rj   rk   r   r   r   r/   �  s(    "




z checkPinCodeVerified_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  | jd k	r�|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr>   r2   r   r3   r	   r�   ru   r   r   r   r'   �  s    


z!checkPinCodeVerified_result.writec                 C   s   d S rF   r   rv   r   r   r   rw   �  s    z$checkPinCodeVerified_result.validatec                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS rx   r   ry   r   r   r   r}   �  s   �z8checkPinCodeVerified_result.__repr__.<locals>.<listcomp>r~   r   r�   r�   r   r   r   r�   �  s    �z$checkPinCodeVerified_result.__repr__c                 C   s   t || j�o| j|jkS rF   r�   r�   r   r   r   r�   �  s    z"checkPinCodeVerified_result.__eq__c                 C   s
   | |k S rF   r   r�   r   r   r   r�   �  s    z"checkPinCodeVerified_result.__ne__)NNr�   r   r   r   r   r>   �  s   
r>   )!�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   �thrift.TRecursiver   �sysrV   �ttypesr   �thrift.transportr   �all_structs�objectr   r   r?   r&   �appendrJ   rh   rc   r1   r�   rU   r=   r�   r>   r�   r   r   r   r   �<module>	   sB   LS8
�D
�8
�D
�